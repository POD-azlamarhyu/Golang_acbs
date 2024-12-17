sendfile="./main.go"
mv_dir="./src/resolved"

echo -e move directory: $mv_dir
read -p "contest: " contest
read -p "question: " ques

mkdir -p "./src/${mv_dir}/${contest}"
mv $sendfile "./src/${mv_dir}/${contest}/${ques}.go"
ga "./src/${mv_dir}/${contest}/${ques}.go"
gc "${mv_dir} ${ques} resolved."

cp_file="./src/template/main.go"

cp  $cp_file $sendfile