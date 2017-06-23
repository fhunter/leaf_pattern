module branch(p1=[[0,1],[1,0],[4,2],[2,4]],ht=1) {
     color("red") linear_extrude(height=ht) polygon(points=p1,convexity=2);
};

module node(p1=[0,0],width=1,ht=1) {
    $fn=10;
    color("lime") translate(p1) cylinder(h=ht,d=width);
};

module growpoint(p1=[0,0]) {
    $fn=10;
    color("yellow") translate(p1) cylinder(h=5,d=1);
};
