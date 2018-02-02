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
leaf.set_size(xsize, ysize)
leaf.set_initial_point_num(pointnum)
leaf.set_max_vein_points(maxveinpoints)
leaf.set_death_distance(deathdistance)
leaf.set_growth_speed(growthspeed)
leaf.set_add_grow_density(addGrowthDensity)
leaf.make_initial_point(Point(0, ysize / 2))
leaf.add_initial_grow_points()  # Add first growpoints
frames.append(leaf.draw())

while True:
    leaf.dispose_grow_points()
    frames.append(leaf.draw())
    if not leaf.add_grow_points():
        break
    if not leaf.develop_veins():
        break

frames[0].save(filename, save_all=True, append_images=frames[1:])
