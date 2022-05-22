package main
import (
    "syscall/js"
)


const (
	width = 400
	height = 400
)

func main() {

    // get canvas object using js syscalls
    var canvas js.Value = js.
    Global().
    Get("document").
    Call("getElementById", "canvas")

    // initialize the canvas
    canvas.Set("height", height)
    canvas.Set("width", width)
    canvas.Set("style", "border: thin solid black")
    // get the rendering context
    ctx := canvas.Call("getContext", "2d")
    ctx.Call("clearRect", 0, 0, width, height)

    //Mandelbrot parameters
    N:=100 // Number of iteration per pixel
    // Simulation space
    xmin := -1.5
    xmax := 0.5
    ymin := -1.0
    ymax := 1.0

    imgData := ctx.Call("createImageData", width, height) // get image data
    data := imgData.Get("data") // get pixel array

    for y := 0; y < height; y++ { // loop over pixels
            for x := 0; x < width; x++ {
                // scaled value of pixel poisition according to our space
                var xx float64 = xmin + (xmax-xmin)*float64(x)/float64(width)
                var yy float64 = ymin +(ymax-ymin)*float64(y)/float64(height)

                index := 4 * (y*width + x) // pixel index in array | 1 pixel = (r,g,b,a)

                // complex number z = r + j*i
                r := 0.0 // real part
                i := 0.0 // imaginary part
                for n := 0; n<N; n++ { // convergence iteration
                    if(r*r+i*i>4.0){ // if does not converge draw color according to the iteration number
                        data.SetIndex(index, int(float64(n)/float64(N)*255.0))
                        data.SetIndex(index+1, int(float64(n)/float64(N)*255.0))
                        data.SetIndex(index+2, int(float64(n)/float64(N)*255.0))
                        data.SetIndex(index+3, 255)
                        break
                    }

                    // iterating the value of z = z*z+c with c = xx + j*yy
                    rt := r*r-i*i + xx
                    i = 2.0*r*i + yy
                    r=rt
                
                }
               

                
                
               
            }
    }
    ctx.Call("putImageData", imgData, 0, 0) // write image data
    println("Finished")


}