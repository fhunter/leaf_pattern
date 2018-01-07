#!/usr/bin/env python
import PIL
import sys
import cmath
import random
import PIL.Image
import PIL.ImageDraw
from datatypes import Tree, Point

filename="leaf1.gif"

if len(sys.argv) >= 2:
	filename = sys.argv[1]

print filename

xsize = 1024
ysize = 1024
pointnum = 10000
maxveinpoints = 40000
deathdistance = 4.0
growthspeed = 1.0
addGrowthDensity = 0.000

im = PIL.Image.new("RGB", (xsize, ysize, ), 0)

def Norm(point):
	return cmath.rect(1, cmath.phase(point))

def distance(point1, point2):
	return abs(point2-point1)

def drawgrowpoints(image, points):
	draw = PIL.ImageDraw.Draw(image)
	#PIL.ImageDraw.ImageDraw.ellipse(xy, fill=None, outline=None)
	for i in points:
		x = i.real
		y = i.imag
		draw.ellipse((x-deathdistance/2,y-deathdistance/2,x+deathdistance/2,y+deathdistance/2),outline=(255,0,0))
	return image

def makeinitialgrowpoints(maxx,maxy,number):
	points=list()
	for i in xrange(0,number):
		point = complex(random.random()*maxx, random.random()*maxy)
		points.append(point)
	return points

def addgrowpoints(points, maxx, maxy, density):
	toadd = int(maxx*maxy*density)
	for i in xrange(0,toadd):
		point = complex(random.random()*maxx, random.random()*maxy)
		points.append(point)
	return points

growpoints = makeinitialgrowpoints(xsize,ysize,pointnum)

im=drawgrowpoints(im, growpoints)

initialpoint=complex(0,ysize/2)
treedata= Tree(initialpoint)

treedata.walk()


#func findClosest(p complex128, vec []complex128) int {
#    if len(vec)==0 {
#        return math.MaxInt32
#    }
#    minDistance:= distance(p,vec[0])
#    minIndex:= 0
#    for i, j:= range vec {
#        if minDistance > distance(p, j) {
#            minIndex = i
#            minDistance = distance(p, j)
#        }
#    }
#    return minIndex
#}
#
#func weight(tree map[int] []int, node int) float64 {
#    var ret float64 = 0.5
#    if node > len(tree) {
#        return ret
#    }
#    if len(tree[node]) == 0 {
#        //This is a leaf
#        ret = 1
#    } else {
#        ret = 0
#        //This is a junction - 
#        for _, k:= range tree[node] {
#            tmp:= weight(tree, k)
#            ret += tmp*tmp
#        }
#        ret= math.Sqrt(ret)
#    }
#    return ret
#}
#
#func calc_weights(tree map[int] []int,length int) []float64 {
#    weights:= make([]float64, length + 1, length+1)
#    for i, _:= range tree {
#        weights[i]= weight(tree, i)
#    }
#    return weights
#}
#
#var frame int = 0
#
#func dumpall(growpoints []complex128, veinNodes []complex128, tree map[int] []int, influence []int) {
#//    dumpall_str(growpoints, veinNodes, tree, influence, "") //This is working dump
#    dumpscad_str(veinNodes, tree, "")
#    str:= fmt.Sprintf("%06d", frame)
#    frame++
#//    dumpall_str(growpoints, veinNodes, tree, influence, str) //This is frame dump
#    dumpscad_str(veinNodes, tree, str)
#}
#
#func dumpall_str(growpoints []complex128, veinNodes []complex128, tree map[int] []int, influence []int, postfix string) {
#    f,err := os.Create("./leaf"+postfix+".dot")
#    check(err)
#    defer f.Close()
#    fmt.Fprintln(f,"graph T {")
#    // grow points output
#    fmt.Fprintln(f,"node [color=\"green\"]")
#    for i, t := range growpoints {
#        //Draw position - xxx [ label = i, pos = "0,0!" ]
#        fmt.Fprint(f,"grownode",i," [label=\"",i,"\", pos=\"",real(t)*32,",",imag(t)*32,"!\" ] ")
#        fmt.Fprintln(f,"")
#    }
#
#    fmt.Fprintln(f,"node [shape=none,color=\"red\", fillcolor=\"transparent\", bgcolor=\"transparent\"]")
#    //Print veinNodes
#    for i, t := range veinNodes {
#//        fmt.Fprint(f,"veinNode",i," [label=\"",i,"\", pos=\"",real(t)*32,",",imag(t)*32,"!\" ] ")
#        fmt.Fprint(f,"veinNode",i," [label=\"\", pos=\"",real(t)*32,",",imag(t)*32,"!\" ] ")
#        fmt.Fprintln(f,"")
#    }
#
#    //Print links
#    fmt.Fprintln(f,"edge [tailclip=false,headclip=false,color=\"red\"]")
#    weights:= calc_weights(tree,len(veinNodes))
#    for i, t:= range tree {
#        for _, k:= range t {
#            fmt.Fprint(f,"veinNode",i,"--veinNode",k,"[penwidth=",int(weights[i]*12+0.5),"]")
#            fmt.Fprintln(f,"")
#        }
#    }
#    fmt.Fprintln(f,"")
#
#    // Print influence
#    fmt.Fprintln(f,"edge [tailclip=true,headclip=false,color=\"blue\",style=\"dotted\"]")
#    for i, t:= range influence {
#        if t<len(veinNodes) {
#            fmt.Fprint(f,"grownode",i,"--veinNode",t)
#            fmt.Fprintln(f,"")
#        }
#    }
#    fmt.Fprintln(f,"}")
#    fmt.Fprintln(f," ")
#}
#
#func dumpscad_str(veinNodes []complex128, tree map[int] []int, postfix string) {
#    f,err := os.Create("./data"+postfix+".scad")
#    check(err)
#    defer f.Close()
#    weights:= calc_weights(tree,len(veinNodes))
#    //Print nodes
#    for i, t:= range veinNodes {
#        fmt.Fprint(f,"node(p1=[",real(t),",",imag(t),"],width=",1,",ht=",(1+weights[i]),");")
#        fmt.Fprintln(f,"")
#    }
#    fmt.Fprintln(f," ")
#    fmt.Fprintln(f," ")
#    //Print links
#    /*
#    for i, t:= range tree {
#        for _, k:= range t {
#            fmt.Fprint(f,"branch(p1=[",real(veinNodes[i]),",",imag(veinNodes[i]),"],p2=[",real(veinNodes[k]),",",imag(veinNodes[k]),"],width1=",weights[i],",width2=",weights[k],");")
#            fmt.Fprintln(f,"")
#        }
#    }
#    */
#    fmt.Fprintln(f," ")
#    fmt.Fprintln(f," ")
#}
#
#
#func main() {
#    //point fill. square 100x100
#    rand.Seed(time.Now().Unix())
#    growpoints := make([]complex128, pointnum, pointnum)
#    growpoints = makeInitialGrowPoints(growpoints, 100.0, 100.0, pointnum)
#
#    veinNodes := make([]complex128, 0, maxveinpoints)
#
#    tree := make(map[int] []int)
#
#    veinNodes = append(veinNodes,50)
#
#    tree[0] = make([]int,0, maxveinpoints)
#
#    for (len(growpoints) > 0) {
#        //Make lists of influence
#        influence := make([]int, len(growpoints), len(growpoints)) //each growpoint is an influence for specific vein point
#        for i, _:= range influence {
#            //Go over all influence points and gather distances. fill the closest
#            influence[i] = findClosest(growpoints[i],veinNodes)
#        }
#        dumpall(growpoints, veinNodes, tree, influence)
#        //Calculate growth vectors
#        {
#            newNodes := make([]complex128, 0, maxveinpoints)
#            for i, _:= range veinNodes {
#                p:= 0+0i //Initial vector
#                needAdd := false
#                for j, k:= range influence {
#                    //Summ vectors
#                    if k == i {
#                        pnt := growpoints[j] - veinNodes[i]
#                        pnt = Norm(pnt)
#                        p += pnt
#                        needAdd = true
#                    }
#                }
#                if needAdd {
#                    p = Norm(p)
#                    p = p * complex128(growthSpeed)
#                    p = veinNodes[i] + p
#                    newNodes = append(newNodes,p)
#                    tree[i] = append(tree[i], len(veinNodes)-1 + len(newNodes))
#                }
#            }
#            for _, t:= range newNodes {
#                veinNodes = append(veinNodes, t)
#            }
#
#        }
#        growpoints = addGrowPoints(growpoints,100.0,100.0,addGrowthDensity)
#        newGrowPoints:= make([]complex128,0,len(growpoints))
#        //Delete any growth points that are too close
#        for _, p:= range growpoints {
#            flag := true
#            for _, t:= range veinNodes {
#                if distance(p, t) < deathdistance {
#                    flag = false
#                    break
#                }
#            }
#            if flag {
#                newGrowPoints = append(newGrowPoints, p)
#            }
#        }
#        fmt.Println("# growpoints",len(growpoints))
#        growpoints = newGrowPoints
#    }
#    dumpall(growpoints, veinNodes, tree, make([]int,0))
#
#
#}
#
#
#

im.save(filename)
