# arrayOperations
Small library for performing union, intersect, difference and distinct operations on arrays in goLang

I don't promise that these are optimized, but they work!

Limited to arrays of strings and uint64 (for now), but the code can be easily changed to accept any type.

## Usage

<dl>
	<dt><h3>1. IntersectString(args...[]string) []string</h3>
		<dd>Find the intersection of two arrays.
		<dd>e.g. a1 = ["1" "2" "2" "4" "6"]; a2 = ["2" "4" "5"]
		<dd>IntersectString(a1, a2) >> ["2" "4"]
	<dt><h3>2. UnionString(args...[]string) []string</h3>
		<dd>Find the union of two arrays.
		<dd>e.g. a1 = ["1" "2" "2" "4" "6"]; a2 = ["2" "4" "5"]
		<dd>UnionString(a1, a2) >> ["1" "2" "4" "5" "6"]
	<dt><h3>3. DifferenceString(args...[]string) []string</h3>
		<dd>Find the union of two arrays.
		<dd>e.g. a1 = ["1" "2" "2" "4" "6"]; a2 = ["2" "4" "5"]
		<dd>DifferenceString(a1, a2) >> ["5" "6"]
	<dt><h3>4. DistinctString(args []string) []string</h3>
		<dd>Remove duplicate values from one array.
		<dd>e.g. a1 = ["1" "2" "2" "4" "6"];
		<dd>DistinctString(a1) >> ["1" "2" "4" "6"]
</dl>

## Other Types
Uint64 functions work the same as string functions, but end in Uint64, e.g. IntersectUint64

Furthermore, a multidimensional array can be used as input by adding "Arr" to the end of the function name e.g. IntersectStringArr(arr [][]string) []string

## Performance

Here's how well it performs in an ubuntu vm on my personal laptop. vm was given 6 cores (3 virtual) @ 2.40 GHz and 8 GB of ram.

**Trials**

| type	| n	| ms	| us / n |
|:--:|:--:|:--:|:--:|
| intersect	 | 10	 | .02	 | 2.32 | 
| intersect	 | 100	 | .36	 | 3.65 | 
| intersect	 | 1,000	 | 2.79	 | 2.79 | 
| intersect	 | 10,000	 | 46.78	 | 4.68 | 
| intersect	 | 100,000	 | 299.61	 | 3 | 
| intersect	 | 1,000,000	 | 3,945.11	 | 3.95 | 
| intersect	 | 10,000,000	 | 44,849.92	 | 4.48 | 
| union	 | 10	 | .04	 | 3.64 | 
| union	 | 100	 | .26	 | 2.65 | 
| union	 | 1,000	 | 2.04	 | 2.04 | 
| union	 | 10,000	 | 14.45	 | 1.44 | 
| union	 | 100,000	 | 119.48	 | 1.19 | 
| union	 | 1,000,000	 | 1,631.15	 | 1.63 | 
| union	 | 10,000,000	 | 16,327.95	 | 1.63 | 
| difference	 | 10	 | .05	 | 4.59 | 
| difference	 | 100	 | .33	 | 3.34 | 
| difference	 | 1,000	 | 3.69	 | 3.69 | 
| difference	 | 10,000	 | 31.86	 | 3.19 | 
| difference	 | 100,000	 | 290.81	 | 2.91 | 
| difference	 | 1,000,000	 | 3,904.97	 | 3.9 | 
| difference	 | 10,000,000	 | 44,719.1	 | 4.47 | 
| distinct	 | 10	 | .01	 | 1.4 | 
| distinct	 | 100	 | .31	 | 3.12 | 
| distinct	 | 1,000	 | 2.3	 | 2.3 | 
| distinct	 | 10,000	 | 29.46	 | 2.95 | 
| distinct	 | 100,000	 | 151.91	 | 1.52 | 
| distinct	 | 1,000,000	 | 1,551.91	 | 1.55 | 
| distinct	 | 10,000,000	 | 15,942.66	 | 1.59 | 


**Averages**

| type	| avg. us / n |
|:--:|:--:|
| union	| 2.032 |
| distinct | 2.062 |
| intersect	| 3.552 |
| difference | 3.727 |

## License
MIT