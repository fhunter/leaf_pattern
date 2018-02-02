import cmath


class Point:
    def __init__(self, x, y):
        self.point = complex(x, y)

    def norm(self):
        return cmath.rect(1, cmath.phase(self.point))

    def distance(self, point2):
        return abs(point2 - self.point)
