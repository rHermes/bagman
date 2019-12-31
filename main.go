package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"

	"github.com/cheggaaa/pb/v3"
	"golang.org/x/net/proxy"
)

var (
	proxyStr   = flag.String("proxy", "socks5://127.0.0.1:9150", "The address of the proxy")
	boardFlag  = flag.String("board", "wg", "The board we wish to read from")
	outdirFlag = flag.String("out", "bag", "The path to put the media files")
)

func fatalf(fmtStr string, args interface{}) {
	fmt.Fprintf(os.Stderr, fmtStr, args)
	os.Exit(-1)
}

func getTorClient() (*http.Client, error) {
	// Create a transport that uses Tor Browser's SocksPort.  If
	// talking to a system tor, this may be an AF_UNIX socket, or
	// 127.0.0.1:9050 instead.
	tbProxyURL, err := url.Parse(*proxyStr)
	if err != nil {
		return nil, err
	}

	// Get a proxy Dialer that will create the connection on our
	// behalf via the SOCKS5 proxy.  Specify the authentication
	// and re-create the dialer/transport/client if tor's
	// IsolateSOCKSAuth is needed.
	tbDialer, err := proxy.FromURL(tbProxyURL, proxy.Direct)
	if err != nil {
		return nil, err
	}

	// Make a http.Transport that uses the proxy dialer, and a
	// http.Client that uses the transport.
	tbTransport := &http.Transport{Dial: tbDialer.Dial}
	client := &http.Client{Transport: tbTransport}

	return client, nil
}

func getMedia(c *http.Client, basedir, board string, thread CatalogThread, post Post) error {

	// first we have to check if the basedir path exists
	filename := fmt.Sprintf("%d%s", post.Tim, post.Ext)
	fdir := filepath.Join(basedir, board, fmt.Sprintf("%d", thread.No))
	fjson := filepath.Join(fdir, "meta.json")
	fpath := filepath.Join(fdir, filename)

	if err := os.MkdirAll(fdir, 0744); err != nil {
		return err
	}

	fj, err := os.OpenFile(fjson, os.O_CREATE|os.O_WRONLY|os.O_EXCL, 0644)
	if err != nil {
		if !os.IsExist(err) {
			return err
		}
	} else {
		jser := json.NewEncoder(fj)
		if err := jser.Encode(thread); err != nil {
			fj.Close()
			os.Remove(fjson)
			return err
		}
		fj.Close()
	}

	fp, err := os.OpenFile(fpath, os.O_CREATE|os.O_WRONLY|os.O_EXCL, 0644)
	if err != nil {
		// If we have it we don't download it
		if os.IsExist(err) {
			return nil
		}
		return err
	}

	u := fmt.Sprintf("https://i.4cdn.org/%s/%s", board, filename)
	resp, err := c.Get(u)
	if err != nil {
		fp.Close()
		os.Remove(fpath)
		return err
	}
	defer resp.Body.Close()

	_, err = io.Copy(fp, resp.Body)
	if err != nil {
		fp.Close()
		os.Remove(fpath)
		return err
	}

	fp.Close()
	return nil
}

func getThread(c *http.Client, board string, threadNum int) (*Thread, error) {
	u := fmt.Sprintf("https://a.4cdn.org/%s/thread/%d.json", board, threadNum)
	resp, err := c.Get(u)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	js := json.NewDecoder(resp.Body)

	var thr Thread
	if err := js.Decode(&thr); err != nil {
		return nil, err
	}

	return &thr, nil

}

func getCatalog(c *http.Client, board string) (Catalog, error) {
	u := fmt.Sprintf("https://a.4cdn.org/%s/catalog.json", board)
	resp, err := c.Get(u)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	js := json.NewDecoder(resp.Body)

	var cat Catalog
	if err := js.Decode(&cat); err != nil {
		return nil, err
	}

	return cat, nil
}

func ripChan(c *http.Client) error {
	// Now we must try
	cat, err := getCatalog(c, *boardFlag)
	if err != nil {
		return err
	}

	for _, b := range cat {
		fmt.Printf("Page: %d\n", b.Page)
		for _, t := range b.Threads {
			fmt.Printf("\tNo: %d\n", t.No)
			fmt.Printf("\t\tSubject: %s\n", t.Sub)

			thr, err := getThread(c, *boardFlag, t.No)
			if err != nil {
				return err
			}
			fmt.Printf("\t\tNum posts: %d\n", len(thr.Posts))

			imgs := []Post{}
			for _, v := range thr.Posts {
				if v.Filename != "" {
					imgs = append(imgs, v)
				}
			}
			fmt.Printf("\t\tNum images: %d vs %d\n", len(imgs), t.Images)
			bar := pb.Full.Start(len(imgs))

			for _, img := range imgs {
				if err := getMedia(c, *outdirFlag, *boardFlag, t, img); err != nil {
					return err
				}
				bar.Increment()
			}
			bar.Finish()
		}
	}

	return nil
}

func main() {
	flag.Parse()
	client, err := getTorClient()
	if err != nil {
		fatalf("Couldn't create tor agent: %s\n", err)
	}

	if err := ripChan(client); err != nil {
		fatalf("Couldn't rip Nyafuu: %#v\n", err)
	}

}
