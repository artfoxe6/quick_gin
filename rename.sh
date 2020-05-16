#!/bin/bash

# 重命名项目名
# ./rename.sh new_name
rename(){
    pwd=`pwd`
    path=${pwd%/*}
    current=${pwd##*/}
    echo $pwd $path $current
    read_dir $pwd $current $1
    mv $pwd $path/$1
}
read_dir(){
    for file in `ls -a $1`
    do
        if [ -d $1"/"$file  ]
        then
            if [[ $file != '.' && $file != '..' && $file != 'apidoc' &&
            $file != 'go.mod' && $file != 'go.sum' && $file != '.idea' && $file != '.git'  ]]
            then
                read_dir $1"/"$file $2 $3
            fi
        else
			      sed -i "s/$2/$3/" $1"/"$file
        fi
    done

}
rename $1

