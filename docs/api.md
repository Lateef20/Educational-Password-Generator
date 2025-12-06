POST /generate

Input: `{ length, useLower, useUpper, useDigits, useSymbols }`

Output: 
```
{
    "password": "generatedValue",
    "entropyBits": 75,
    "stats": {
    "charsetSize": 62,
    "crackTimeOnline": "20 years",
    "crackTimeOffline": "3 days"
    },
    "contentKeys": ["entropy_basics", "rate_limiting_explainer"]
}
```


POST /analyze

Input: `{ "password": "userInput" }`

Output: 
```
{
    "entropyBits": 40,
    "flags": ["too_short", "only_letters"],
    "crackTimeOffline": "2 hours"
    "contentKeys": ["short_password_warning", "character_set_tip"]
}
```

