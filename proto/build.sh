#!/usr/bin/env bash

usage()
{
    cat << EOF
    use: bash build.sh [-h] [-v] [dir...]
        -h 显示帮助信息
        -v 只检查语法，不编译
        dir 目录名，不传参的情况下递归处理执行目录下的所有文件夹，也可以传一个或者多个定向检查或生成
EOF
    exit 0
}

function listFiles()
{
    if [ ! -f "$1/prototool.yaml" ]; then
        for file in `ls $1`;
        do
                if [ -d "$1/$file" ]; then
                    ff="$1/$file"
                    if [ -f "$1/$file/prototool.yaml" ]; then
                        (lintAndGenerate "$1/$file" $2)
                    else
                        listFiles "$1/$file" $2
                    fi
                fi
        done
    else
        (lintAndGenerate $1 $2)
    fi
}

function lintAndGenerate()
{
    pwd=`pwd`
    dir=$1
    cd $dir
    if [ ! -d $dir ]; then
        echo "$dir not exists"
        exit -1
    fi

    if [ ! -f "$dir/prototool.yaml" ]; then
        echo "$dir not a prototool.yaml"
        exit -1
    fi

    out=$(prototool lint)
    if [ $? -ne 0 ]; then
        echo -e "\033[40;37m lint \033[0m $dir \033[31m FAIL \033[0m"
        echo $out
        exit -1
    fi
    echo -e "\033[40;37m lint \033[0m $dir \033[32m SUCCESS \033[0m"

    if [ "$2" == "false" ]; then
        out=$(prototool all)
        if [ $? -ne 0 ]; then
            echo -e "\033[40;37m generate \033[0m $dir \033[31m FAIL \033[0m"
            echo $out
            exit -1
        fi
        echo -e "\033[40;37m generate \033[0m $dir \033[32m SUCCESS \033[0m"
    fi
    exit $?
}


#检查参数
checkOption="false"
while getopts 'hv' OPT; do
    case $OPT in
        v) checkOption="true";;
        h) usage
            exit 0
            ;;
        ?) usage
            exit 0
            ;;
    esac
done

#遍历余下参数并获取所有传入的目录
shift $(($OPTIND - 1))
declare -i index
index=0
declare -a arr
while true; do
    if [ ! $1 ]; then
        break
    else
        if [ -d $1 ]; then

            arr[$index]=$(cd $1;pwd)
        else
            echo "$1 目录不存在"
            exit -1
        fi
        index=$index+1
    fi
    shift
done

#删除旧生成的pb文件
work_path=$(dirname $(readlink -f "$0"))
echo $work_path../go/errors
rm -rf $work_path/../go/errors
rm -rf $work_path/../go/proto
rm -rf $work_path/../php/UtoExceptions
rm -rf $work_path/../php/UtoProto
rm -rf $work_path/../php/GPBMetadata

#默认不传以当前目录为起点
if [ ${#arr[@]} -eq 0 ]; then
    arr[0]=`pwd`
fi

#遍历所有目录并执行操作
for dir in ${arr[*]}
do
    listFiles $dir $checkOption
done

