# speech2text
I made my own hmm speech to text as much as possible.

## description
Refer to [Julius](https://julius.osdn.jp/index.php).
speech to text with monophone.

## required files
- HMM definition file
  - monophone
- Word N-gram language model file
  - ARPA canonical format
- Word dictionary file
  - Use a format that is almost equivalent to the HTK dictionary format


## Used repository created by myself
- [github.com/yut-kt/gowave](https://github.com/yut-kt/gowave)
  - Wave file read support for Go language
- [github.com/yut-kt/goft](https://github.com/yut-kt/goft)
  - Fourier Transformation support for Go language
- [github.com/yut-kt/gowindow](https://github.com/yut-kt/gowindow)
  - Window Function support for Go language
- [github.com/yut-kt/gomfcc](https://github.com/yut-kt/gomfcc)
  - MFCC(Mel Frequency Cepstral Coefficient) support for golang.
- [github.com/yut-kt/gostruct](https://github.com/yut-kt/gostruct)
  - Convenient golang struct