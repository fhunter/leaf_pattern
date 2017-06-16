# leaf_pattern
Leaf pattern generator

This is initial project to learn Go language

## Visualisation

### Png files

~~~
for i in *.dot;do
    dot -Kneato -n2 -Tpng -Gdpi=30 -o `basename ${i} .dot`.png $i
done

~~~

### Video

~~~
ffmpeg -r 5 -i leaf00%4d.png -c:a copy -c:v libx264 -crf 18 -preset veryslow myvideo.mpg
~~~
