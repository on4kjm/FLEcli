# FLE File specification

## General

* line comment: "#" => "`^#`s" 
* Multi line comment: enclosed in "{}"

## Header

* **mycall**	The mycall keyword is your station call sign and it is mandatory. This is the logging station's call sign (the call sign used over the air). Saving ADIF files without your own call is prevented.
* **operator** 	The call sign of the operator can optionally be added to the ADIF file, it is not mandatory. Use the operator keyword and add the call sign of the operator. It can be different from the station call sign but it may not be the name of the operator! The operator keyword can be repeated further down the log with other operator call signs if they change.
* **qslmsg**	QSL Messages valid for the entire log can be entered in the header section of the log. It is not mandatory.
* **mywwff**	The mywwff keyword is used to register your own WWFF reference number in WWFF Logging. See also chapter WWFF logging. The syntax is: AAFF-CCCC: AA = national prefix, CCCC = 4-digit numeric code (e.g. ONFF-0001).
* **mysota**	The mysota keyword is used to register your own SOTA reference number in SOTA Logging. The syntax is: AA/NN-CCC: Association/Name-3-digit numeric Code (e.g. G/CE-001). Your own SOTA reference number is mandatory for SOTA Logging.
* **nickname**	The nickname keyword can be used for eQSL ADIF uploads. See chapter Uploading logs to eQSL.cc.
* **date**	The date format is year-month-day (YYYY-MM-DD), e.g. 2016-12-31. Year, month and day may be abbreviated and you may use separators other than dash.

## validations
* call 