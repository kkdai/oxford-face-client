package main

import (
	"bufio"
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	. "github.com/kkdai/oxford-face"
	"github.com/spf13/cobra"
)

type DetectedFace struct {
	FaceID string
}

var StoreFaces []DetectedFace

func Init() {
}

//Remove file extension for renaming process
func removeExtension(full string) string {
	extension := filepath.Ext(full)
	return full[0 : len(full)-len(extension)]
}

//Draw a rectangle on src source when detect a face on it
func drawRectange(srcFile string, r image.Rectangle) error {
	//Check src file
	fileSrc, err := os.Open(srcFile)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer fileSrc.Close()
	imgSrc, _ := jpeg.Decode(fileSrc)

	fN := removeExtension(srcFile)
	newFileName := fmt.Sprintf("%s_update.%s", fN, filepath.Ext(srcFile))
	fileDst, err := os.Create(newFileName)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer fileDst.Close()

	jpg := image.NewRGBA(image.Rect(r.Min.X, r.Min.Y, r.Max.X, r.Max.Y))
	draw.Draw(jpg, imgSrc.Bounds().Add(image.Pt(10, 10)), imgSrc, imgSrc.Bounds().Min, draw.Src)
	jpeg.Encode(fileDst, jpg, nil)
	return nil
}

func toggleLogging(enable bool) {
	if enable {
		log.SetOutput(os.Stdout)
	} else {
		log.SetOutput(ioutil.Discard)
	}
}

func printConsole() {
	fmt.Println("Command:( C:Create S:Subscription P:Publish R:RemoveTopic V:Verbose G:Read Q:exit )")
	fmt.Printf(":>")
}

func main() {
	var key string
	var verbose bool

	rootCmd := &cobra.Command{
		Use:   "oxford-face-client",
		Short: "Client to connect to Project Oxford Face API services",
		Run: func(ccmd *cobra.Command, args []string) {

			toggleLogging(verbose)
			if key == "" {
				fmt.Println("Input key is empty, make sure you input key with -k or key")
				return
			}

			client := NewFace(key)
			if client == nil {
				fmt.Println("Cannot connect to server, please check your setting.")
				return
			}
			quit := false
			scanner := bufio.NewScanner(os.Stdin)
			printConsole()
			for !quit {

				var param1, param2 string

				if !scanner.Scan() {
					break
				}
				line := scanner.Text()
				parts := strings.Split(line, " ")
				cmd := parts[0]
				if len(parts) > 1 {
					param1 = parts[1]
				}
				if len(parts) > 2 {
					param2 = parts[2]
				}

				fmt.Println(param1, param2)
				//var err error
				switch cmd {
				case "a", "A", "Add", "add", "ADD": //CREATE face
					var res []byte
					var errRsp *ErrorResponse
					if param1 != "" {
						if strings.Contains(param1, "http") {
							res, errRsp = client.DetectUrl(nil, param1)
						} else {
							res, errRsp = client.DetectFile(nil, param1)
						}
					} else {
						//Not valid input for param1
						printConsole()
						continue
					}

					if errRsp.Err != nil {
						fmt.Println("Err:", errRsp.Err)
						continue
					}

					gotFaces := NewFaceResponse(res)
					if gotFaces == nil {
						fmt.Println("Got error on result :", string(res))
						continue
					}

					for _, gotFace := range gotFaces {
						fmt.Println("New Face:", gotFace)
						newFace := DetectedFace{}
						newFace.FaceID = gotFace.Faceid
						StoreFaces = append(StoreFaces, newFace)
					}

				case "Q", "q":
					quit = true
				case "V", "v":
					verbose = !verbose
					toggleLogging(verbose)
					fmt.Println("Switch verbose to ", verbose)
				case "l", "L", "list", "LIST":
					for index, face := range StoreFaces {
						fmt.Println(index, ":", face)
					}
				default:
					fmt.Println("Command not support.")
				}

				if quit != true {
					printConsole()
				}
			}
		},
	}
	rootCmd.Flags().StringVarP(&key, "key", "k", "", "Project Oxford key, please export your key to environment first, `export MSFT_KEY=12234`")
	rootCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Verbose")
	rootCmd.Execute()
}
