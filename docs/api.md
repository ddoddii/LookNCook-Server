# API Docs

## Recipe

- Example Response for Recipe
```json
{
  "ingredients": [
    {
      "name": "Burrata Cheese",
      "locationDescription": "On the middle shelf, on the left"
    },
    {
      "name": "Butter",
      "locationDescription": "On the middle shelf, on the left"
    },
    {
      "name": "Cream Cheese",
      "locationDescription": "On the middle shelf, on the left"
    },
    {
      "name": "Yogurt",
      "locationDescription": "On the middle shelf, on the left"
    },
    {
      "name": "Milk",
      "locationDescription": "On the middle shelf, on the right"
    },
    {
      "name": "Eggs",
      "locationDescription": "On the bottom shelf, on the right"
    },
    {
      "name": "Tupperware with food",
      "locationDescription": "On the bottom shelf, in the middle"
    },
    {
      "name": "Spam",
      "locationDescription": "On the top shelf, on the right"
    },
    {
      "name": "Jar of something",
      "locationDescription": "On the top shelf, in the middle"
    },
    {
      "name": "Green onions",
      "locationDescription": "On the top shelf, on the left"
    },
    {
      "name": "Flower-decorated pitcher with a purple liquid",
      "locationDescription": "On the middle shelf, on the right"
    },
    {
      "name": "Bread",
      "locationDescription": "On the middle shelf, in the middle"
    }
  ],
  "recipes": [
    {
      "title": "Scrambled Eggs and Toast",
      "difficulty": "Easy",
      "ingredients": [
        {
          "name": "Eggs",
          "locationDescription": "On the bottom shelf, on the right"
        },
        {
          "name": "Milk",
          "locationDescription": "On the middle shelf, on the right"
        },
        {
          "name": "Butter",
          "locationDescription": "On the middle shelf, on the left"
        },
        {
          "name": "Bread",
          "locationDescription": "On the middle shelf, in the middle"
        }
      ],
      "instructions": [
        {
          "step": "Crack the eggs into a bowl and whisk them together with a fork."
        },
        {
          "step": "Heat a pan over medium heat and add a knob of butter."
        },
        {
          "step": "Pour the egg mixture into the pan and let it cook for 2-3 minutes, stirring constantly."
        },
        {
          "step": "Season the eggs with salt and pepper to taste."
        },
        {
          "step": "Toast the bread in a toaster or on a griddle."
        },
        {
          "step": "Serve the scrambled eggs on the toast and enjoy!"
        }
      ]
    },
    {
      "title": "Cheese and Yogurt Parfait",
      "difficulty": "Easy",
      "ingredients": [
        {
          "name": "Yogurt",
          "locationDescription": "On the middle shelf, on the left"
        },
        {
          "name": "Burrata Cheese",
          "locationDescription": "On the middle shelf, on the left"
        },
        {
          "name": "Green onions",
          "locationDescription": "On the top shelf, on the left"
        },
        {
          "name": "Tupperware with food",
          "locationDescription": "On the bottom shelf, in the middle"
        }
      ],
      "instructions": [
        {
          "step": "Layer 1/2 cup of yogurt, 1/2 cup of burrata cheese, and 1/4 cup of chopped green onions in a parfait glass or jar."
        },
        {
          "step": "Repeat the layers until the glass is full."
        },
        {
          "step": "Top with a dollop of yogurt and a sprinkle of chopped green onions."
        },
        {
          "step": "Serve immediately or chill for later."
        }
      ]
    },
    {
      "title": "Creamy Spinach Pasta",
      "difficulty": "Moderate",
      "ingredients": [
        {
          "name": "Cream Cheese",
          "locationDescription": "On the middle shelf, on the left"
        },
        {
          "name": "Spinach",
          "locationDescription": "Tupperware with food, On the bottom shelf, in the middle"
        },
        {
          "name": "Milk",
          "locationDescription": "On the middle shelf, on the right"
        },
        {
          "name": "Butter",
          "locationDescription": "On the middle shelf, on the left"
        },
        {
          "name": "Pasta",
          "locationDescription": "Not mentioned"
        },
        {
          "name": "Salt and pepper",
          "locationDescription": "Not mentioned"
        }
      ],
      "instructions": [
        {
          "step": "Cook the pasta according to the package instructions."
        },
        {
          "step": "While the pasta is cooking, melt the butter in a large skillet over medium heat."
        },
        {
          "step": "Add the spinach and cook until wilted."
        },
        {
          "step": "Stir in the cream cheese and milk until melted and smooth."
        },
        {
          "step": "Season with salt and pepper to taste."
        },
        {
          "step": "Drain the pasta and add it to the skillet with the sauce."
        },
        {
          "step": "Toss to coat the pasta in the sauce and serve."
        }
      ]
    }
  ]
}
```

## Kitchen Environment

- Example Response for kitchen hazard environment

```json
{
  "hazards": [
    {
      "type": "utensil",
      "content": "There is a chef holding a knife in the kitchen. The handle of the knife is located at the bottom right corner of the image."
    },
    {
      "type": "hazard",
      "content": "The knife is sharp and could cut the chef if they are not careful."
    },
    {
      "type": "utensil",
      "content": "There is a hot stove in the kitchen. The stove is located in the middle of the image."
    },
    {
      "type": "hazard",
      "content": "The stove is hot and could burn the chef if they are not careful."
    },
    {
      "type": "utensil",
      "content": "There is a pot on the stove. The pot is located in the middle of the image."
    },
    {
      "type": "hazard",
      "content": "The pot is hot and could burn the chef if they are not careful."
    }
  ]
}
```