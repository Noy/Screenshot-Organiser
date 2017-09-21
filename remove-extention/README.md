### I added this in here as I thought it matched the topic well, you can customise this to anything, but in my case I had about 300 directories all ending with '.com, .new, .info' etc. and I needed to rename all of them to 'directorynamecom' instead of 'directoryname.com'

#### You could always change it to something like:

```golang
for _, file := range files {
                    // You can do file name checks like I did.
                
                    // Here you can customise the name to something else
		            os.Rename(file.Name(), "some string.txt") // Or add more complex logic.
}
```

#### I have the logic for random strings in the [Image-Extension](https://github.com/Noy/Screenshot-Organiser/blob/master/image-extension/main.go#L31) class if you want to use that.

## Hope this fits your needs!
