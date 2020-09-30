package writetxt

import "io"

func Copy(r io.Reader, w io.Writer) (n int, err error) {
    var wrote int
    b := make([]byte, 512)

    for {
        _ , err = r.Read(b)
        if err != nil && err != io.EOF {
            return
        }

        wrote, err = w.Write(b)
        n += wrote
        if err != nil {
            return
        }
    }
}
