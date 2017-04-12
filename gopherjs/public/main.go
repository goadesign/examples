package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	_ "strconv"

	"github.com/goadesign/examples/upload/client"

	"github.com/gopherjs/gopherjs/js"
)

var (
	c   *client.Client
	ctx context.Context
	doc *js.Object
)

func _goForIt() {

	req, err := c.NewUploadImageRequest(ctx, "/api/images")
	if nil != err {
		//TODO
		return
	}

	buf := new(bytes.Buffer)
	mw := multipart.NewWriter(buf)

	jsFile := doc.Call("getElementById", "myFile").Get("files").Index(0)    // Get the file object
	filePart, err := mw.CreateFormFile("file", jsFile.Get("name").String()) //Set the name of the filename
	if nil != err {
		fmt.Printf("Failed to create the image Part with error %s\n", err)
		return
	}

	b := blobToBytes(jsFile) // Convert to blob
	filePart.Write(b)
	if nil != err {
		fmt.Printf("Failed to copy with error %s\n", err)
		return
	}
	req.Header.Set("Content-Type", mw.FormDataContentType())
	mw.Close()
	req.Body = ioutil.NopCloser(buf)
	req.ContentLength = int64(buf.Len())

	resp, err := c.Client.Do(ctx, req)
	if nil != err {
		fmt.Printf("Failed to get response %s\n", err)
	}
	var text string
	if resp.StatusCode == http.StatusOK {
		image, _ := c.DecodeImageMedia(resp)
		text = fmt.Sprintf("Successfully upload image with new id %d", image.ID)
		doc.Call("getElementById", "myImage").Set("value", image.ID)
		_showImage()
	} else {
		errResponse, _ := ioutil.ReadAll(resp.Body)
		text = fmt.Sprintf("Failed to upload image with error: %s", string(errResponse))
	}
	displayResponse("upload_result", text)

}

// blobToBytes converts a Blob to []byte.
func blobToBytes(blob *js.Object) []byte {
	var b = make(chan []byte)
	fileReader := js.Global.Get("FileReader").New()
	fileReader.Set("onload", func() {
		b <- js.Global.Get("Uint8Array").New(fileReader.Get("result")).Interface().([]byte)
	})
	fileReader.Call("readAsArrayBuffer", blob)
	return <-b
}

func displayResponse(id, text string) {

	resDiv := doc.Call("getElementById", id)
	resDiv.Set("innerHTML", text)
}

func goForIt() {
	go func() {
		_goForIt()
	}()
}

func _showImage() {
	image := doc.Call("getElementById", "myImage").Get("value")
	resp, err := c.ShowImage(ctx, client.ShowImagePath(image.Int()))

	var text string
	if nil != err {
		text = err.Error()
	} else if resp.StatusCode == http.StatusOK {
		image, _ := c.DecodeImageMedia(resp)
		buff := new(bytes.Buffer)
		json.NewEncoder(buff).Encode(image)
		text = fmt.Sprintf("Image: <pre>%s</pre><img src=\"%s\"></img>", buff.String(), fmt.Sprintf("/download/%d", image.ID))
	} else {
		errResponse, _ := ioutil.ReadAll(resp.Body)
		text = fmt.Sprintf("Failed to get image with error: %s", string(errResponse))
	}
	displayResponse("show_result", text)
}

func showImage() {
	go func() {
		_showImage()
	}()
}

func onLoad() {

	uploadButton := doc.Call("getElementById", "uploadButton")
	uploadButton.Call("addEventListener", "click", goForIt)
	showButton := doc.Call("getElementById", "showImage")
	showButton.Call("addEventListener", "click", showImage)
}

func main() {
	c = client.New(nil)
	c.Host = "127.0.0.1:8080"
	ctx = context.Background()
	doc = js.Global.Get("document")
	js.Global.Set("onload", onLoad)
}
