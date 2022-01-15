package design_pattern

// 迪米特法则：搜索引擎抓取网页地功能
// 1.1 网页请求
type HtmlRequest struct {
	address string
	data    []byte
}
type Html struct{}

// 1.2 负责底层网络通信，根据请求获取数据
type NetworkTransporter struct{}

// 注意1：此处依赖HtmlRequest
func (n NetworkTransporter) send(request *HtmlRequest) Html {
	//....
	return Html{}
}

// 1.3 通过URL获取网页
// 注意2：HtmlDownloader依赖NetworkTransporter
type HtmlDownloader struct {
	transporter NetworkTransporter
}

func (h *HtmlDownloader) downloadHtml(url string) Html {
	html := h.transporter.send(new(HtmlRequest))
	//....
	return html
}

// 1.4 表示网页文档，后续地网页内容抽取/分词/索引都是以此为处理对象
type Document struct {
	html Html
	url  string
}

func NewDocument(url string) *Document {
	return &Document{url: url}
}

// 注意3：Document依赖HtmlDownloader
func (d *Document) getDocument() {
	downloader := new(HtmlDownloader)
	d.html = downloader.downloadHtml(d.url)
}

func (n NetworkTransporter) send2(address string, data []byte) Html {
	//....
	return Html{}
}
func (h *HtmlDownloader) downloadHtml2(url string) Html {
	request := new(HtmlRequest)
	html := h.transporter.send2(request.address, request.data)
	//....
	return html
}

// 2.2 针对Document依赖HtmlDownloader，可以通过简单工厂实现
type DocumentFactory struct {
	downloader HtmlDownloader
}

func (d DocumentFactory) NewDocFactoty(downloader HtmlDownloader) DocumentFactory {
	return DocumentFactory{
		downloader: downloader,
	}
}

type DocumentInf interface {
	getDocument()
}

func (d DocumentFactory) createDocument(url string) DocumentInf {
	html := d.downloader.downloadHtml2(url)
	return &Document{
		url:  url,
		html: html,
	}
}
