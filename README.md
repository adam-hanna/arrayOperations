# arrayOperations
Small library for performing union, intersect, difference and distinct operations on arrays in goLang

I don't promise that these are optimized, but they work!

Limited to arrays of strings, but the code can be easily changed to accept any type.

## Performance
**Trials**

| type	| n	| ms	| us / n |
|:--:|:--:|:--:|:--:|
| intersect	 | 10	 | 0.02	 | 2.32 | 
| intersect	 | 100	 | 0.36	 | 3.65 | 
| intersect	 | 1000	 | 2.79	 | 2.79 | 
| intersect	 | 10000	 | 46.78	 | 4.68 | 
| intersect	 | 100000	 | 299.61	 | 3 | 
| intersect	 | 1000000	 | 3945.11	 | 3.95 | 
| intersect	 | 10000000	 | 44849.92	 | 4.48 | 
| union	 | 10	 | 0.04	 | 3.64 | 
| union	 | 100	 | 0.26	 | 2.65 | 
| union	 | 1000	 | 2.04	 | 2.04 | 
| union	 | 10000	 | 14.45	 | 1.44 | 
| union	 | 100000	 | 119.48	 | 1.19 | 
| union	 | 1000000	 | 1631.15	 | 1.63 | 
| union	 | 10000000	 | 16327.95	 | 1.63 | 
| difference	 | 10	 | 0.05	 | 4.59 | 
| difference	 | 100	 | 0.33	 | 3.34 | 
| difference	 | 1000	 | 3.69	 | 3.69 | 
| difference	 | 10000	 | 31.86	 | 3.19 | 
| difference	 | 100000	 | 290.81	 | 2.91 | 
| difference	 | 1000000	 | 3904.97	 | 3.9 | 
| difference	 | 10000000	 | 44719.1	 | 4.47 | 
| distinct	 | 10	 | 0.01	 | 1.4 | 
| distinct	 | 100	 | 0.31	 | 3.12 | 
| distinct	 | 1000	 | 2.3	 | 2.3 | 
| distinct	 | 10000	 | 29.46	 | 2.95 | 
| distinct	 | 100000	 | 151.91	 | 1.52 | 
| distinct	 | 1000000	 | 1551.91	 | 1.55 | 
| distinct	 | 10000000	 | 15942.66	 | 1.59 | 

**Averages**

| type	| avg. us / n |
|:--:|:--:|
| union	| 2.032 |
| distinct | 2.062 |
| intersect	| 3.552 |
| difference | 3.727 |

## License
MIT