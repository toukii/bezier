# bezier

Bézier curve(贝塞尔曲线) in svg
=============

```
go get github.com/toukii/bezier
cd $GOPATH/src/github.com/toukii/bezier/bezier_in_svg
go test -v
```

![No smooth bezier](http://7xku3c.com1.z0.glb.clouddn.com/github/no-smooth.svg?v=0.1)

![Smooth bezier](http://7xku3c.com1.z0.glb.clouddn.com/github/smooth.svg)


```go
NewPoint(110, 105),
NewPoint(220, 240),
NewPoint(130, 250),
NewPoint(180, 350),
NewPoint(280, 450),
NewPoint(480, 150),
NewPoint(111, 211),
NewPoint(222, 122),
NewPoint(333, 433),
NewPoint(444, 344),
NewPoint(555, 655),
NewPoint(666, 566),
NewPoint(777, 877),
NewPoint(888, 788),
NewPoint(999, 999),
```

svg path:

version1:
```svg
M110 105 C220 236, 220 244, 175 245M175 245 C131 245, 129 255, 155 300M155 300 C175 344, 185 356, 230 400M230 400 C258 465, 302 435, 380 300M380 300 C483 155, 477 145, 295 180M295 180 C116 211, 106 211, 166 166M166 166 C204 104, 240 140, 277 277M277 277 C315 415, 351 451, 388 388M388 388 C426 326, 462 362, 499 499M499 499 C537 637, 573 673, 610 610M610 610 C648 548, 684 584, 721 721M721 721 C759 859, 795 895, 832 832M832 832 C869 778, 907 798, 999 999
```
![](http://7xku3c.com1.z0.glb.clouddn.com/github/bezier.svg?v=0.2)

vsersion2:
```
<path d="M110 105 C220 238, 220 241, 175 245M175 245 C130 248, 130 253, 155 300M155 300 C178 348, 182 353, 230 400M230 400 C274 454, 295 440, 380 300M380 300 C481 152, 479 148, 295 180M295 180 C114 211, 110 211, 166 166M166 166 C217 117, 234 134, 277 277M277 277 C321 421, 338 438, 388 388M388 388 C439 339, 456 356, 499 499M499 499 C543 643, 560 660, 610 610M610 610 C661 561, 678 578, 721 721M721 721 C765 865, 782 882, 832 832M832 832 C881 785, 899 794, 999 999" stroke="#ff6140" stroke-width="3" fill="none"></path>
```


![](http://7xku3c.com1.z0.glb.clouddn.com/github/bezier-smooth.svg?v=0.3)
