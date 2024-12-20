sendfile="./main.go"
mv_dir="./src/resolved"

read -p "contest th. : " number 
read -p "question of atcoder beginners contest : " ques

contest="abc"+number

echo -e move directory: "./src/${mv_dir}/${contest}"

mkdir -p "./src/${mv_dir}/${contest}"
mv $sendfile "./src/${mv_dir}/${contest}/${ques}.go"
git add ./src/$mv_dir/$contest/$ques.go
git commit -m "${mv_dir} ${ques} resolved."

cp_file="./src/template/main.go"

cp  $cp_file $sendfile