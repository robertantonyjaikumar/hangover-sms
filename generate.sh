#!/bin/bash


# Prompt the user for their name
# echo "Please enter your name:"
# read name



# Input string
file_name="user_group"

# Set IFS to a space character
IFS='_' read -r -a array <<< "$file_name"

title_case=""

# Print the array
i=0
for item in "${array[@]}"; do
    
    title_case=$(echo "$item")
    echo "$item"
    ((i=i+1))
done

echo "$i"

echo "$title_case"

# echo "Original: $input_word"
# echo "Lowercase: $lowercase"
# echo "Uppercase: $uppercase"
# echo "Titlecase: $titlecase"
# echo "Snake_case: $snake_case"
# echo "Kebab-case: $kebab_case"
# echo "PascalCase: $pascal_case"
# echo "UPPER_SNAKE_CASE: $upper_snake_case"


# cp models/a_user_group.go "models/$name.go"

#sleep 2

# sed -i '' "s/UserGroup/$model_name/g" "models/$name.go"
# sed -i '' "s/user_groups/$table_name/g" "models/$name.go"
# sed -i '' "s/user group/$logger_name/g" "models/$name.go"
# sed -i '' "s/userGroup/$variable_name/g" "models/$name.go"


#cp routes/auth.go "routes/$name.go"

#cp controllers/auth.go "controllers/$name.go"

# echo "Hello, $name!"