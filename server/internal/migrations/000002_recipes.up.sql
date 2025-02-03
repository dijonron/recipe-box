CREATE TABLE IF NOT EXISTS rs_recipes (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(), 
    name VARCHAR(255) NOT NULL,   
    chef VARCHAR(255),             
    cookbook VARCHAR(255),         
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,   
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
);

CREATE TABLE IF NOT EXISTS rs_recipe_ingredients (
    ingredient_id UUID PRIMARY KEY NOT NULL,                     
    recipe_id UUID NOT NULL,                          
    amount FLOAT NOT NULL,                            
    measurement INT NOT NULL,                         
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,   
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,   
    deleted_at TIMESTAMP NULL,
    FOREIGN KEY (recipe_id) REFERENCES rs_recipes(id)   
);