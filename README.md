# annodate

`annodate` prepends/appends day of week to stdin or file

# Installation

```
go get https://github.com/grepinsight/annodate
```

# Usages
```
$ annodate -h

Usage of annodate:
  -d string
    	Delimiter (default " ")
  -e	Add column at the end?
  -f int
    	Field that contains date data (default 1)
```

# Examples
```
$ cat testFile.txt # e.g. hledger activity
2017-12-11 **************************
2017-12-12 ********
2017-12-13 **********
2017-12-14 **************
2017-12-15 ****************
2017-12-16 ****

$ cat testFile.txt | annodate 
Monday 2017-12-11 **************************
Tuesday 2017-12-12 ********
Wednesday 2017-12-13 **********
Thursday 2017-12-14 **************
Friday 2017-12-15 ****************
Saturday 2017-12-16 ****

Or 

$ annodate testFile.txt
Monday 2017-12-11 **************************
Tuesday 2017-12-12 ********
Wednesday 2017-12-13 **********
Thursday 2017-12-14 **************
Friday 2017-12-15 ****************
Saturday 2017-12-16 ****

$cat testFile2.csv  
**************************,2017-12-11
********,2017-12-12
**********,2017-12-13
**************,2017-12-14
****************,2017-12-15
****,2017-12-16
********,2017-12-17
**********,2017-12-18

$cat testFile2.csv | annodate -d , -f 2 -e
**************************,2017-12-11,Monday
********,2017-12-12,Tuesday
**********,2017-12-13,Wednesday
**************,2017-12-14,Thursday
****************,2017-12-15,Friday
****,2017-12-16,Saturday
********,2017-12-17,Sunday
**********,2017-12-18,Monday

```


