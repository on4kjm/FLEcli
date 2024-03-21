# FLE File specification

## General

* line comment: "#" => "`^#`s"
* Multi line comment: enclosed in "{}"

## Header

* **mycall** The mycall keyword is your station call sign and it is mandatory. This is the logging station's call sign (the call sign used over the air). Saving ADIF files without your own call is prevented.
* **operator** The call sign of the operator can optionally be added to the ADIF file, it is not mandatory. Use the operator keyword and add the call sign of the operator. It can be different from the station call sign but it may not be the name of the operator! The operator keyword can be repeated further down the log with other operator call signs if they change.
* **qslmsg** QSL Messages valid for the entire log can be entered in the header section of the log. It is not mandatory.
* **mywwff** The mywwff keyword is used to register your own WWFF reference number in WWFF Logging. See also chapter WWFF logging. The syntax is: AAFF-CCCC: AA = national prefix, CCCC = 4-digit numeric code (e.g. ONFF-0001).
* **mysota** The mysota keyword is used to register your own SOTA reference number in SOTA Logging. The syntax is: AA/NN-CCC: Association/Name-3-digit numeric Code (e.g. G/CE-001). Your own SOTA reference number is mandatory for SOTA Logging.
* **mypota** This registers your POTA reference for POTA logging. POTA is (March 2024) transitioning their references to use ISO country codes. The syntax is AA-##### where AA is an ISO country code and ##### is a 4 to 5 digit number.
* **nickname** The nickname keyword can be used for eQSL ADIF uploads. See chapter Uploading logs to eQSL.cc.
* **date** The date format is year-month-day (YYYY-MM-DD), e.g. 2016-12-31. Year, month and day may be abbreviated and you may use separators other than dash.
* **mycounty** This is field will populate the <MY_CNTY> tag in your ADIF. It may contain spaces. (e.g. mycounty WY,Laramie County)
* **mylat** This field takes your decimal degrees lattitude and populates the adif <MY_LAT> tag.
* **mylon** This field takes your decimal degrees longitutde and populates the adif <MY_LON> tag.

## Validations

* call
* mypota - Any combination of values 0-9 or A-Z one to five times, a dash, and 0-9 four to five times. e.g. US-0609
* mylat - ±90.0 degrees inclusive.
* mylon - ±180.0 degrees inclusive.

## POTA Example

```
# Header
mycall AA0AAA/7
operator AA0AAA
qslmsg Please QSL via LotW
nickname Bamforth National Wildlife Refuge
mypota US-0609
mygrid DN71ci
mylat 41.3724
mylon -105.752
mycounty WY,Albany County

# QSOs
date 2024-03-20
10m 28.028 cw
1234 w1aw 599 599 @Bob <Wow! W1AW!>
1235 n0aw 599 599 US-0001 <My first P2P!>
1237 ab1wx 599 599
1300 w0ny 319 339 <Tough copy>

15m 21.043 ssb
1311 n9fz 55 55
1333 kj4ase 59 59
```