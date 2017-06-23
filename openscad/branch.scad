module branch(p1=[0,0],p2=[1,1],width1=1.5,width2=1,ht=1) {
    distance=sqrt((p2.x-p1.x)*(p2.x-p1.x)+(p2.y-p1.y)*(p2.y-p1.y));
    angle = atan((p2.x-p1.x)/(p2.y-p1.y));
    //echo(distance);
    //echo(angle);
    p1_1=[p1.x+width1*sin(angle)/2,p1.y-cos(angle)*width1/2];
    p1_2=[p1.x-width1*sin(angle)/2,p1.y+cos(angle)*width1/2];
    p2_1=[p2.x+width2*sin(angle)/2,p2.y-cos(angle)*width2/2];
    p2_2=[p2.x-width2*sin(angle)/2,p2.y+cos(angle)*width2/2];
    //echo(p1,p1_1,p2_1,p2,p2_2,p1_2);
//    union() {
        linear_extrude(height=ht) polygon(points=[p1,p1_1,p2_1,p2,p2_2,p1_2],convexity=1);
//        $fn=10;
//        translate(p1) cylinder(height=ht,d=width1);
//        translate(p2) cylinder(height=ht,d=width2);
//    }
};

module node(p1=[0,0],width=1,ht=1) {
    $fn=10;
    color("lime") translate(p1) cylinder(h=ht,d=width);
};

module growpoint(p1=[0,0]) {
    $fn=10;
    color("yellow") translate(p1) cylinder(h=5,d=1);
};
