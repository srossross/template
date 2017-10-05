package lib

import (
  "os"
  "strings"
  "io/ioutil"
  "path"
)

type InputIteratorResult struct {
	In, Out string
  Error error
}

func InputIterator(args []string, recursive bool) <-chan InputIteratorResult {
    ch := make(chan InputIteratorResult)
    go func() {
        for _, name := range args {

          var In, Out string
          InOut := strings.SplitN(name,":", 2)
          if len(InOut) > 1 {
            In, Out = InOut[0], InOut[1]
          } else {
            In = InOut[0]
            Out = InOut[0]
          }

          fi, err := os.Stat(In)
          if err != nil {
            ch <- InputIteratorResult{name, "", err}
            return;
          }

          switch mode := fi.Mode(); {
          case mode.IsDir():
              // do directory stuff
              InFiles, err := ioutil.ReadDir(In)
              if err != nil {
                ch <- InputIteratorResult{name, "", err}
                return;
              }

              if Out != "-" {
                err = os.MkdirAll(Out, os.ModePerm)
                if err != nil {
                  ch <- InputIteratorResult{name, "", err}
                  return
                }
              }

              for _, InF := range InFiles {
                OutPutPath := Out
                if Out != "-" {
                  OutPutPath = path.Join(Out, InF.Name())
                }
                ch <- InputIteratorResult{path.Join(In, InF.Name()), OutPutPath, err}
              }

          case mode.IsRegular():
              // do file stuff
              ch <- InputIteratorResult{In, Out, err}
          }

        }
        close(ch) // Remember to close or the loop never ends!
    }()
    return ch
}
