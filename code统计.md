1. xxx代码量统计 
git log --author="xxx" --pretty=tformat: --numstat | awk '{ add += $1; subs += $2; loc += $1 - $2 } END { printf "added lines: %s, removed lines: %s, total lines: %s\n", add, subs, loc }'
2. 项目总代码量统计 
git log --author="" --pretty=tformat: --numstat | awk '{ add += $1; subs += $2; loc += $1 - $2 } END { printf "added lines: %s, removed lines: %s, total lines: %s\n", add, subs, loc }'
3. 每个人的统计 
git log --format='%aN' | sort -u | while read name; do echo -en "$name"; git log --author="$name" --pretty=tformat: --numstat | awk '{ add += $1; subs += $2; loc += $1 - $2 } END { printf " added lines: %s, removed lines: %s, total lines: %s\n", add, subs, loc }' -; done
4. cloc统计 
D:/software/tool/cloc ./ --exclude-dir=node_modules,vendor,wkhtmltopdf,wkhtmltopdf0.12.4,resources/assets/library/js,public/dist,storage --exclude-ext=json,jpg,png,xml,cls,doc,docx
