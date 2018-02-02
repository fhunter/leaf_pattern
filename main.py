#!/usr/bin/env python
# vim: set fileencoding=utf-8 :
import sys
from Leaf import Leaf
from Point import Point

filename = "leaf1.gif"

if len(sys.argv) >= 2:
    filename = sys.argv[1]

print filename

frames = []

# Алгоритм:
# 1. Залить в лист ограничительный контур
# 2. Задать начальные точки роста
# 3.

xsize = 1024
ysize = 1024
pointnum = 10000
maxveinpoints = 40000
deathdistance = 4.0
growthspeed = 1.0
addGrowthDensity = 0.000

leaf = Leaf()  # We created a leaf
leaf.setSize(xsize, ysize)
leaf.setInitialPointNum(pointnum)
leaf.setMaxVeinPoints(maxveinpoints)
leaf.setDeathDistance(deathdistance)
leaf.setGrowthSpeed(growthspeed)
leaf.setAddGrowDensity(addGrowthDensity)
leaf.makeinitialPoint(Point(0, ysize / 2))
leaf.addinitialgrowpoints()  # Add first growpoints
frames.append(leaf.draw())

while True:
    leaf.disposegrowpoints()
    frames.append(leaf.draw())
    if not leaf.addgrowpoints():
        break
    if not leaf.developVeins():
        break

frames[0].save(filename, save_all=True, append_images=frames[1:])
