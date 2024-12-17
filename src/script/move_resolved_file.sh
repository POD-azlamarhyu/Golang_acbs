sendfile="./main.go"
mv_dir="resolved"

echo -e move directory: $mv_dir
read -p "question: " ques

mkdir -p ./src/$mv_dir
mv $sendfile "./src/${mv_dir}/${ques}.go"
git add ./src/$mv_dir/$ques.go
git commit -m "${mv_dir} ${ques} resolved."

cp_file="./src/template/main.go"

cp  $cp_file $sendfile