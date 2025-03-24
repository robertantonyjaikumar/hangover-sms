#!/bin/bash

# Prompt the user for the module name
echo "Please enter module name (ex: user, user_group, user_account_type):"
read file_name

# Define source placeholders
sourceFileName="todo"
srcModelName="Todo"
srcTableName="todos"
srcLoggerName="todos-logger"
srcVariableName="vartodo"

# Initialize variables
migrationPattern="// crud-generator-migration"
seedPattern="// crud-generator-seeds"
routePattern="// crud-generator-router"


table_name=""
model_name=""
variable_name=""
logger_name=""

# Split the file_name into parts based on underscores
IFS='_' read -r -a array <<< "$file_name"

# Initialize indices for the loop
i=0
for item in "${array[@]}"; do
    title=$(echo "$item" | awk '{print toupper(substr($0,1,1)) tolower(substr($0,2))}')

    # First item
    if [ $i -eq 0 ]; then
        logger_name="$item"
        model_name="$title"
        variable_name="$item"
        table_name="$item"
    else
        logger_name="$logger_name $item"
        model_name="$model_name$title"
        variable_name="$variable_name$title"
        table_name="$table_name"_"$item"
    fi
    ((i=i+1))
done

# Append 's' to table_name for pluralization
table_name="${table_name}s"

# Prepare content for migration, seed, and route
migration=$(echo "\&$model_name{},\n$migrationPattern")
seedContent=$(echo "\/\/ {Model: \&[]$model_name{}, FileName: \"$table_name.json\", CreateFunc: Seed$model_name},\n$seedPattern")
routeContent=$(echo $variable_name"Group := v1.Group(\"$file_name\")\n{\n$model_name""Routes("$variable_name"Group)\n}\n \n$routePattern")


# Display the generated content (for debugging purposes)
# echo "Route Content: $routeContent"
# echo "Seed Content: $seedContent"

# Copy files to new locations
cp "models/$sourceFileName.go" "models/$file_name.go"
cp "controllers/$sourceFileName.go" "controllers/$file_name.go"
cp "routes/$sourceFileName.go" "routes/$file_name.go"
cp "seeds/$srcTableName.json" "seeds/$table_name.json"
sleep 2



# Replace placeholders in the new files
replace_in_file() {
    local src="$1"
    local dest="$2"
    local pattern="$3"
    local replacement="$4"

    sed -i '' "s|$pattern|$replacement|g" "$dest"
}

# Replace patterns in models, controllers, and routes
replace_in_file "$srcModelName" "models/$file_name.go" "$srcModelName" "$model_name"
replace_in_file "$srcTableName" "models/$file_name.go" "$srcTableName" "$table_name"
replace_in_file "$srcLoggerName" "models/$file_name.go" "$srcLoggerName" "$logger_name"
replace_in_file "$srcVariableName" "models/$file_name.go" "$srcVariableName" "$variable_name"

replace_in_file "$srcModelName" "controllers/$file_name.go" "$srcModelName" "$model_name"
replace_in_file "$srcVariableName" "controllers/$file_name.go" "$srcVariableName" "$variable_name"

replace_in_file "$srcModelName" "routes/$file_name.go" "$srcModelName" "$model_name"

# Replace route, migration, and seed patterns in relevant files
replace_in_file "$routePattern" "routes/routers.go" "$routePattern" "$routeContent"
replace_in_file "$migrationPattern" "models/migrator.go" "$migrationPattern" "$migration"
replace_in_file "$seedPattern" "models/migrator.go" "$seedPattern" "$seedContent"

# Completion message
echo "Hello, $file_name!"
