package main

// Boards represents the output of the boards api endpoint
type Boards struct {
	Boards []struct {
		Board           string `json:"board"`
		Title           string `json:"title"`
		WsBoard         int    `json:"ws_board"`
		PerPage         int    `json:"per_page"`
		Pages           int    `json:"pages"`
		MaxFilesize     int    `json:"max_filesize"`
		MaxWebmFilesize int    `json:"max_webm_filesize"`
		MaxCommentChars int    `json:"max_comment_chars"`
		MaxWebmDuration int    `json:"max_webm_duration"`
		BumpLimit       int    `json:"bump_limit"`
		ImageLimit      int    `json:"image_limit"`
		Cooldowns       struct {
			Threads int `json:"threads"`
			Replies int `json:"replies"`
			Images  int `json:"images"`
		} `json:"cooldowns"`
		MetaDescription string `json:"meta_description"`
		IsArchived      int    `json:"is_archived,omitempty"`
		Spoilers        int    `json:"spoilers,omitempty"`
		CustomSpoilers  int    `json:"custom_spoilers,omitempty"`
		ForcedAnon      int    `json:"forced_anon,omitempty"`
		UserIds         int    `json:"user_ids,omitempty"`
		CountryFlags    int    `json:"country_flags,omitempty"`
		CodeTags        int    `json:"code_tags,omitempty"`
		WebmAudio       int    `json:"webm_audio,omitempty"`
		MinImageWidth   int    `json:"min_image_width,omitempty"`
		MinImageHeight  int    `json:"min_image_height,omitempty"`
		Oekaki          int    `json:"oekaki,omitempty"`
		SjisTags        int    `json:"sjis_tags,omitempty"`
		TextOnly        int    `json:"text_only,omitempty"`
		RequireSubject  int    `json:"require_subject,omitempty"`
		TrollFlags      int    `json:"troll_flags,omitempty"`
		MathTags        int    `json:"math_tags,omitempty"`
	} `json:"boards"`
	TrollFlags map[string]string `json:"troll_flags"`
}

// CatalogThread is a thread from the catalog page
type CatalogThread struct {
	No            int    `json:"no"`
	Sticky        int    `json:"sticky,omitempty"`
	Closed        int    `json:"closed,omitempty"`
	Now           string `json:"now"`
	Name          string `json:"name"`
	Com           string `json:"com,omitempty"`
	Filename      string `json:"filename"`
	Ext           string `json:"ext"`
	W             int    `json:"w"`
	H             int    `json:"h"`
	TnW           int    `json:"tn_w"`
	TnH           int    `json:"tn_h"`
	Tim           int64  `json:"tim"`
	Time          int    `json:"time"`
	Md5           string `json:"md5"`
	Fsize         int    `json:"fsize"`
	Resto         int    `json:"resto"`
	Capcode       string `json:"capcode,omitempty"`
	SemanticURL   string `json:"semantic_url"`
	Replies       int    `json:"replies"`
	Images        int    `json:"images"`
	LastModified  int    `json:"last_modified"`
	Sub           string `json:"sub,omitempty"`
	Bumplimit     int    `json:"bumplimit,omitempty"`
	Imagelimit    int    `json:"imagelimit,omitempty"`
	OmittedPosts  int    `json:"omitted_posts,omitempty"`
	OmittedImages int    `json:"omitted_images,omitempty"`
	LastReplies   []struct {
		No       int    `json:"no"`
		Now      string `json:"now"`
		Name     string `json:"name"`
		Com      string `json:"com"`
		Filename string `json:"filename,omitempty"`
		Ext      string `json:"ext,omitempty"`
		W        int    `json:"w,omitempty"`
		H        int    `json:"h,omitempty"`
		TnW      int    `json:"tn_w,omitempty"`
		TnH      int    `json:"tn_h,omitempty"`
		Tim      int64  `json:"tim,omitempty"`
		Time     int    `json:"time"`
		Md5      string `json:"md5,omitempty"`
		Fsize    int    `json:"fsize,omitempty"`
		Resto    int    `json:"resto"`
	} `json:"last_replies,omitempty"`
}

// Catalog represents the output of the catalog endpoint
type Catalog []struct {
	Page    int             `json:"page"`
	Threads []CatalogThread `json:"threads"`
}

// Post is a struct for holding the posttype
type Post struct {
	No          int    `json:"no"`
	Now         string `json:"now"`
	Name        string `json:"name"`
	Sub         string `json:"sub,omitempty"`
	Com         string `json:"com,omitempty"`
	Filename    string `json:"filename,omitempty"`
	Ext         string `json:"ext,omitempty"`
	W           int    `json:"w,omitempty"`
	H           int    `json:"h,omitempty"`
	TnW         int    `json:"tn_w,omitempty"`
	TnH         int    `json:"tn_h,omitempty"`
	Tim         int64  `json:"tim,omitempty"`
	Time        int    `json:"time"`
	Md5         string `json:"md5,omitempty"`
	Fsize       int    `json:"fsize,omitempty"`
	Resto       int    `json:"resto"`
	Bumplimit   int    `json:"bumplimit,omitempty"`
	Imagelimit  int    `json:"imagelimit,omitempty"`
	SemanticURL string `json:"semantic_url,omitempty"`
	Replies     int    `json:"replies,omitempty"`
	Images      int    `json:"images,omitempty"`
	UniqueIps   int    `json:"unique_ips,omitempty"`
}

// Thread describes a thread
type Thread struct {
	Posts []Post `json:"posts"`
}
