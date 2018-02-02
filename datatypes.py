#!/usr/bin/env python
import cmath
import PIL
import PIL.Image
import PIL.ImageDraw
import random

class Tree:
    def __init__(this, root):
        this.element=root
        this.branches=list()
    def walk(this):
        for i in this.branches:
            i.walk()
        this.action()
    def addbranch(this,element):
        this.branches.append(element)
    def action(this):
        print this.element

class Point:
    def __init__(this, x, y):
        this.point = complex(x,y)
    def norm(this):
        return cmath.rect(1, cmath.phase(this.point))
    def distance(this, point2):
        return abs(point2-this.point)

class Leaf:
    sizex=0
    sizey=0
    ninitialpoints = 0
    maxVeintPoints = 0
    numVeinPoints = 0
    deathdistance = 0.0
    growthspeed = 0.0
    addGrowDensity = 0.0

    tree = None
    growpoints = list()

    def __init__(this):
        pass
    def setSize(this, x, y):
        this.sizex = x
        this.sizey = y
    def setInitialPointNum(this,  n):
        this.ninitialpoints = n
    def setMaxVeinPoints(this, n):
        this.maxVeintPoints = n
    def setDeathDistance(this,  distance):
        this.deathdistance = distance
    def setGrowthSpeed(this,  speed):
        this.growthspeed = speed
    def setAddGrowDensity(this, density):
        this.addGrowthDensity = density
    def makeinitialPoint(this, pnt):
        this.tree = Tree(pnt)
    def addinitialgrowpoints(this):
        points = list()
        for i in xrange(0, this.ninitialpoints):
            point = complex(random.random()*this.sizex, random.random()*this.sizey)
            points.append(point)
        this.growpoints = points
        return len(points)
    def addgrowpoints(this):
        "Return true if adding was successful. False - if there was no free space to add points"
        return True # TODO: this is unimplemented yet
        toadd = int(this.sizex*this.sizey*this.addGrowthDensity)
        points = this.growpoints
        for i in xrange(0,toadd):
            point = complex(random.random()*this.sizex, random.random()*this.sizey)
            points.append(point)
        this.growpoints = points
    def disposegrowpoints(this):
        "TODO: this is not implemented yet. Should delete any grow points that are too close to the vein nodes"
        pass
    def developVeins(this):
        "Return true if operation succeeded. Return false, if there is nothing to develop, or if tree size exceeded"
        return False

    def draw(this):

        """ This function draws the current state of the leaf - veins, attractions, and grow points"""
        image = PIL.Image.new("RGB", (this.sizex, this.sizey, ), 0) # Initial image
        draw = PIL.ImageDraw.Draw(image)
        deathdistance = this.deathdistance
        #here goes grow point drawing
        for i in this.growpoints:
            x = i.real
            y = i.imag
            draw.ellipse((x-deathdistance/2,y-deathdistance/2,x+deathdistance/2,y+deathdistance/2),outline=(255,0,0))
        #here goes venation drawing
        #here goes vein connection drawing
        #here goes attractions drawing
        return image


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
