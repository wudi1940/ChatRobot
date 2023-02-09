package processer

//
//import (
//	"encoding/binary"
//	"fmt"
//	"github.com/gordonklaus/portaudio"
//	"os"
//	"os/signal"
//	"strconv"
//	"strings"
//	"time"
//)
//
//type AudioInput struct {
//	BasePath string
//	FileName string
//}
//
//func (a *AudioInput) Record() {
//
//	if strings.EqualFold(a.BasePath, "") {
//		a.BasePath = "audio"
//	}
//	fmt.Println("Recording.  Press Ctrl-C to stop.")
//	a.FileName = a.BasePath + "/" + strconv.FormatInt(time.Now().Unix(), 10)
//	if !strings.HasSuffix(a.FileName, ".aiff") {
//		a.FileName += ".aiff"
//	}
//	f, err := os.Create(a.FileName)
//	if err != nil {
//		panic(err)
//	}
//
//	sig := make(chan os.Signal, 1)
//	signal.Notify(sig, os.Interrupt, os.Kill)
//	nSamples := 0
//
//	portaudio.Initialize()
//	defer portaudio.Terminate()
//	in := make([]int32, 64)
//	stream, err := portaudio.OpenDefaultStream(1, 0, 44100, len(in), in)
//	if err != nil {
//		panic(err)
//	}
//	defer stream.Close()
//
//	err = stream.Start()
//	if err != nil {
//		panic(err)
//	}
//
//	for {
//		err = stream.Read()
//		if err != nil {
//			panic(err)
//		}
//
//		err = binary.Write(f, binary.BigEndian, in)
//		if err != nil {
//			panic(err)
//		}
//
//		nSamples += len(in)
//		select {
//		case <-sig:
//			return
//		default:
//		}
//	}
//
//	err = stream.Stop()
//	if err != nil {
//		panic(err)
//	}
//
//}
