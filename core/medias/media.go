package medias

type Media struct {
	ID        string `json:"media_id"`
	Hash      string `json:"media_hash"`
	Timestamp int64  `json:"created_at"`
	MediaType int    `json:"media_type"`
	FileSize  int    `json:"file_size"`
	SrcWidth  int    `json:"src_width"`
	SrcHeight int    `json:"src_height"`
	Flags     []int  `json:"media_flags"`
	Version   int    `json:"media_version"`
	Variants  map[string]*Media
	IpfsHash  string `json:"ipfs_hash"`
	UpdatedOn int64  `json:"updated_on"`
}
