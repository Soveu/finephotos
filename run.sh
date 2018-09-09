#!/bin/bash

TEMPDIR=/tmp/finephotos
WALLDIR=~/wall

mkdir -p $TEMPDIR
mkdir -p $WALLDIR

./main > $TEMPDIR/urls
wget -i $TEMPDIR/urls -P $TEMPDIR

cd $TEMPDIR

for photo in $(ls)
do
	width=$(identify -format "%w" $photo)
	height=$(identify -format "%h" $photo)

	if [ $width -ge 1000 ] && [ $height -ge 800 ]; then
		mv $photo $WALLDIR
	fi
done

