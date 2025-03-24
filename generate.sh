#!/bin/bash


# Prompt the user for their name
echo "Please enter module name(ex: user, user_group, user_account_type):"
read file_name


# Set IFS to a space character
IFS='_' read -r -a array <<< "$file_name"

logger_name=""
model_name=""
variable_name=""
table_name=""

# Print the array
i=0
for item in "${array[@]}"; do
    title=$(echo "$item" | awk '{print toupper(substr($0,1,1)) tolower(substr($0,2))}')
    if [ $i -eq 0 ] ; then
        logger_name=$(echo "$item")
        model_name=$(echo "$title")
        variable_name=$(echo "$item")
        table_name=$(echo "$item")
    else
        logger_name=$(echo "$logger_name $item")
        model_name=$(echo "$model_name$title")
        variable_name=$(echo "$variable_name$title")
        table_name=$(echo $table_name"_"$item)
    fi
    ((i=i+1))
done

# echo "$logger_name"
# echo "$model_name"
# echo "$table_name"

table_name=$(echo "$table_name"'s')

migrationPattern=$(echo "// crud-generator-migration")
seedPattern=$(echo "// crud-generator-seeds")
routePattern=$(echo "// crud-generator-router")

migration=$(echo  "\&$model_name{},\n $migrationPattern")
seedContent=$(echo "\/\/ {Model: \&[]$model_name{}, FileName: \"$table_name.json\", CreateFunc: Seed$model_name},\n $seedPattern")
routeContent=$(echo $variable_name"Group := v1.Group(\"$file_name\")\n{\n"$model_name"Routes("$variable_name"Group)\n}\n \n $routePattern")

echo $routeContent
echo $seedContent


cp models/todo.go "models/$file_name.go"
cp controllers/todo.go "controllers/$file_name.go"
cp routes/todo.go "routes/$file_name.go"
cp seeds/todos.json "seeds/$table_name.json"
sleep 2

sed -i '' "s/Todo/$model_name/g" "models/$file_name.go"
# sed -i '' "s/todos/$table_name/g" "models/$file_name.go"
sed -i '' "s/todos-logger/$logger_name/g" "models/$file_name.go"
sed -i '' "s/vartodo/$variable_name/g" "models/$file_name.go"

sed -i '' "s/Todo/$model_name/g" "controllers/$file_name.go"
sed -i '' "s/todo/$variable_name/g" "controllers/$file_name.go"
sed -i '' "s/Todo/$model_name/g" "routes/$file_name.go"

sed -i '' "s/Todo/$model_name/g" "routes/$file_name.go"

sed -i '' "s|$routePattern|$routeContent|g" "routes/routers.go"
sed -i '' "s/$migrationPattern/$model_name/g" "routes/$file_name.go"
sed -i '' "s|$migrationPattern|$migration|g" "models/migrator.go"
sed -i '' "s|$seedPattern|$seedContent|g" "models/migrator.go"



echo "Hello, $file_name!"