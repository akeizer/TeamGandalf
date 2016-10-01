# Script for generating training / test images
# Uses ImageMagick (convert)

# generally want small images for not-high input dimensionality
isize=20
isizestr=20x20

for shape in triangle square circle
do
	for i in {1..10}
	do
		# square of some size
		let "size = 4 + $RANDOM % 12"
		let "x1 = $RANDOM % (isize - size + 1)"
		let "y1 = $RANDOM % (isize - size + 1)"
		let "x2 = x1 + size"
		let "y2 = y1 + size"

		if [ $shape = "triangle" ]; then
			# pick 3 points in the rectangle area
			# segregate them a bit to make sure it has some area
			let "tx1 = $RANDOM % (size/4) + x1 + (size - size/4)"
			let "ty1 = $RANDOM % size + y1"
			let "tx2 = $RANDOM % (size/2) + x1"
			let "ty2 = $RANDOM % (size/4) + y1 + (size - size/4)"
			let "tx3 = $RANDOM % (size/2) + x1"
			let "ty3 = $RANDOM % (size/4) + y1"
			draw="polygon $tx1,$ty1 $tx2,$ty2 $tx3,$ty3"
		elif [ $shape = "square" ]; then
			draw="rectangle $x1,$y1 $x2,$y2"
		else
			let "cx = (x1 + x2) / 2"
			let "cy = (y1 + y2) / 2"
			# use cx for both points to ensure circle is smaller than the allocated square
			draw="circle $cx,$cy $cx,$y1"
		fi

		eval "convert -size $isizestr canvas:white -stroke black -strokewidth 1 -fill black -draw '$draw' $shape-$i.png"
	done
done
