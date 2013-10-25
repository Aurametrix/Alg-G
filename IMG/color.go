package main
// package color implements a basic color library.
import (
        "image"
        "image/color"
        "image/png"
        "log"
        "os"
)

func main() {
        width, height := 128, 128
        m := image.NewRGBA(image.Rect(0, 0, width, height))
        drawColor(*m)
        outFilename := "color.png"
        outFile, err := os.Create(outFilename)
        if err != nil {
                log.Fatal(err)
        }
        defer outFile.Close()
        log.Print("Saving image to: ", outFilename)
        png.Encode(outFile, m)
}

func drawColor(m image.RGBA) {
        size := m.Bounds().Size()
        for x := 0; x < size.X; x++ {
                for y := 0; y < size.Y; y++ {
// Func.RGBA - (r, g, b, a uint32) 
                       color := color.RGBA{
                                204,
                                204,
                                255,
                                255}
                        m.Set(x, y, color)
                }
        }
}
