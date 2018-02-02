class Tree:
    def __init__(self, root):
        self.element = root
        self.branches = list()

    def walk(self):
        for i in self.branches:
            i.walk()
        self.action()

    def addbranch(self, element):
        self.branches.append(element)

    def action(self):
        print self.element
