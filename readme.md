# Look-N-Cook-Server

This repository is go server for **look-n-cook** project, **google solution challenge 2024**

## What this is made of?


This server if made with `golang` and `echo framework`.

## How does it work?


When client sends a photo of the fridge to the server,
the server generates the location of each food in the fridge and recipes user can make with these food with gemini api.
Before the user starts cooking, user takes a picture of the cooking environment.
The server figures out any kind of hazard there are in the cooking environment.

This server calls gemini api and sends the results to the client.

![_lookncook-server](https://github.com/ddoddii/ddoddii.github.io/assets/95014836/89cc3e3f-15fc-4a5a-bdd2-a0943c2cfa9f)

## API reference

| using                                      | method | url                      | body        | response               |
|--------------------------------------------|--------|--------------------------|-------------|------------------------|
| check server is alive                      | GET    | '/'                       | -           | Hello World            |
| analyze fridge content and generate recipe | POST   | '/GetFridgeRecipe'       | image(jpeg) | example response below |
| analyze any hazard in cooking environment  | POST   | '/GetKitchenEnvironment' | image(jpeg) | example response below |

<details>
<summary> '/GetFridgeRecipe' example response </summary>


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

</details>

<details>
<summary> '/GetKitchenEnvironment' example response </summary>

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
</details>

