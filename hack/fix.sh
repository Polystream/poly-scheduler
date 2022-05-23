dos2unix -n $1 tmp_file
rm $1
mv -f tmp_file $1