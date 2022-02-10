#!/bin/bash


# query_type - (Day | Month | Year) for querying data.

START="2022-02-08 00:00:00"
STOP="2022-02-08 23:59:59"

export TOKEN="$(jq -r .result_data.token login.response)"
export USERID="$(jq -r .result_data.user_id login.response)"
export EMAIL="$(jq -r .result_data.email login.response)"
export APPKEY="$(jq -r .appkey login.post)"
export PSID="$(jq --tab '.result_data.pageList[0].ps_id' getPsList.response)"

export START_TIMESTAMP="$(date -j -f '%Y-%m-%d %H:%M:%S' "${START}" +'%Y%m%d%H%M00')"
export END_TIMESTAMP="$(date -j -f '%Y-%m-%d %H:%M:%S' "${STOP}" +'%Y%m%d%H%M00')"
export START_EPOCH="$(date -j -f '%Y-%m-%d %H:%M:%S' "${START}" +'%s000')"
export END_EPOCH="$(date -j -f '%Y-%m-%d %H:%M:%S' "${STOP}" +'%s000')"
export START_DATE="$(date -j -f '%Y-%m-%d %H:%M:%S' "${START}" +'%Y-%m-%d')"
export END_DATE="$(date -j -f '%Y-%m-%d %H:%M:%S' "${STOP}" +'%Y-%m-%d')"
export START_TIME="$(date -j -f '%Y-%m-%d %H:%M:%S' "${START}" +'%H:%M:%S')"
export END_TIME="$(date -j -f '%Y-%m-%d %H:%M:%S' "${STOP}" +'%H:%M:%S')"

export SERIAL_NO="B2192301528"
# SERIAL_NO="A2192703899" # Inverter sn
export PS_KEY="1129147_11_0_0"
export TASK_ID="$2"

export USER_AGENT="$(cat user_agent.data)"

(cat<<EOF
{
	"user_id": "${USERID}",
	"valid_flag": "1,3",
	"lang": "_en_US",
	"token": "${TOKEN}",
	"appkey": "${APPKEY}",
	"sys_code": "200"
}
EOF
) > default.json
export JSON="$(jq -r . default.json)"

API="$1"
export POSTFILE="${API}.post"
export RESPONSEFILE="${API}.response"
shift

case ${API} in
	'login')
		URL="foo"
		;;

	'getPowerDevicePointNames')
		URL="/v1/reportService/getPowerDevicePointNames"
		# DEVICE_TYPE = 1, 3, 4, 5, 7, 11, 14, 17
		export DEVICE_TYPE="$1"
		if [ "${DEVICE_TYPE}" == "" ]; then DEVICE_TYPE="1"; fi

		RESPONSEFILE="${API}-${DEVICE_TYPE}.response"
		JSON="$(jq -r '. += {"device_type": env.DEVICE_TYPE}' default.json)"
		;;

	# getTemplateList -> queryUserCurveTemplateData -> getPsList -> queryMutiPointDataList
	'queryMutiPointDataList')
		URL="/v1/commonService/queryMutiPointDataList"

		export POINTS="$2"
		if [ "${POINTS}" == "" ]; then POINTS="p83033,p83022"; fi
		# p13116,p13199,p13174,p13028,p13173,p13147,p13112,p13122,p83002,p83032,p83011,p83549,p83072,p83119,p83097,p83102,p83033,p83022,p83006

		export PS_KEY="$3"
		if [ "${PS_KEY}" == "" ]; then PS_KEY="1129147_11_0_0,1129147_11_0_0"; fi
		# 1129147_14_1_1,1129147_14_1_1,1129147_14_1_1,1129147_14_1_1,1129147_14_1_1,1129147_14_1_1,1129147_14_1_1,1129147_14_1_1,1129147_11_0_0,1129147_11_0_0,1129147_11_0_0,1129147_11_0_0,1129147_11_0_0,1129147_11_0_0,1129147_11_0_0,1129147_11_0_0,1129147_11_0_0,1129147_11_0_0,1129147_11_0_0

		export MINUTE_INTERVAL="5"

		export START_TIMESTAMP="${1}000000"
		export END_TIMESTAMP="${1}235900"

		RESPONSEFILE="${API}-${1}.response"
		JSON="$(jq -r '. += {"start_time_stamp": env.START_TIMESTAMP, "end_time_stamp": env.END_TIMESTAMP, "ps_key": env.PS_KEY, "points": env.POINTS, "minute_interval": env.MINUTE_INTERVAL}' default.json)"
		;;

	'getHouseholdStoragePsReport')
		URL="/v1/powerStationService/getHouseholdStoragePsReport"
		# Day
		export DATE_TYPE="1"; export DATE_ID="20220208"
		# Month
		#export DATE_TYPE="2"; export DATE_ID="202202"
		# Year
		#export DATE_TYPE="3"; export DATE_ID="2022"
		# Total
		#export DATE_TYPE="3"; export DATE_ID="2022"
		export DATE_TYPE="$1"
		export DATE_ID="$2"
		if [ "${DATE_TYPE}" == "" ]; then DATE_TYPE="1"; fi
		if [ "${DATE_ID}" == "" ]; then DATE_ID="$(date +'%Y%m%d')"; fi

		RESPONSEFILE="${API}-${DATE_ID}.response"
		JSON="$(jq -r '. += {"ps_id": env.PSID, "date_type": env.DATE_TYPE, "date_id": env.DATE_ID}' default.json)"
		;;


	################################################################################
	'getAPIServiceInfo')
		URL="/v1/commonService/getAPIServiceInfo"
		;;

	'getAccessedPermission')
		URL="/v1/commonService/getAccessedPermission"
		;;

	'getAllDeviceByPsId')
		URL="/v1/devService/getAllDeviceByPsId"
		;;

	'getAllPowerDeviceSetName')
		URL="/v1/devService/getAllPowerDeviceSetName"
		JSON="$(jq -r '.sys_code = "900"' default.json)"
		;;

	'getAllPowerRobotViewInfoByPsId')
		URL="/v1/devService/getAllPowerRobotViewInfoByPsId"
		;;

	'getAllPsIdByOrgIds')
		URL="/v1/devService/getAllPsIdByOrgIds"
		;;

	'getPsReport')
		URL="/v1/reportService/getPsReport"
		;;

	'getPsDetail')
		URL="/v1/powerStationService/getPsDetail"
		;;

	'getPsDetailForSinglePage')
		URL="/v1/powerStationService/getPsDetailForSinglePage"
		;;

	'getPsHealthState')
		URL="/v1/powerStationService/getPsHealthState"
		;;

	'getPowerStationData')
		URL="/v1/powerStationService/getPowerStationData"
		;;

	'getPowerStationBasicInfo')
		URL="/v1/powerStationService/getPowerStationBasicInfo"
		;;

	'getPowerStationInfo')
		URL="/v1/powerStationService/getPowerStationInfo"
		;;

	'getPowerStationPR')
		URL="/v1/powerStationService/getPowerStationPR"
		;;

	'getPsListStaticData')
		URL="/v1/powerStationService/getPsListStaticData"
		;;

	'getReportData')
		URL="/v1/powerStationService/getReportData"
		;;

	'getReportExportColumns')
		URL="/v1/reportService/getReportExportColumns"
		;;

	'getBoxData')
		URL="/v1/devService/getBoxData"
		;;

	'getPowerTrendDayData')
		URL="/v1/powerStationService/getPowerTrendDayData"
		JSON="$(jq -r '. += {"beginTime": env.START_TIME, "endTime": env.END_TIME, "ps_id": env.PSID}' default.json)"
		;;

	'getPowerTrendMonthData')
		URL="/v1/powerStationService/getPowerTrendMonthData"
		;;

	'getPowerTrendYearData')
		URL="/v1/powerStationService/getPowerTrendYearData"
		;;


	################################################################################
	# Invalid data.
	'getDataFromHBase')
		URL="/v1/commonService/getDataFromHBase"
		export TABLE_NAME="inverter"
		export PRIMARY_KEY="inverter"
		export POINT_IDS="81012"
		JSON="$(jq -r '. += {"table": env.TABLE_NAME, "primaryKey": env.PRIMARY_KEY, "pointIds": env.POINT_IDS, "datetime": env.START_TIMESTAMP}' default.json)"
		;;

	# Invalid args.
	'getDataFromHbaseByRowKey')
		URL="/v1/commonService/getDataFromHbaseByRowKey"
		export TABLE_NAME="inverter"
		export ROW_KEY_LIST="inverter"
		export FAMILY_NAME="Hellstrom"
		export COLUMN="2"
		JSON="$(jq -r '. += {"table_name": env.TABLE_NAME, "row_key_list": env.ROW_KEY_LIST, "family_name": env.FAMILY_NAME, "column": env.COLUMN}' default.json)"
		;;

	# Unknown args.
	'getListMiFromHBase')
		URL="/v1/commonService/getListMiFromHBase"
		export TABLE_NAME="inverter"
		export PRIMARY_KEY="inverter"
		JSON="$(jq -r '. += {"table": env.TABLE_NAME, "primaryKey": env.PRIMARY_KEY}' default.json)"
		;;

	# Invalid data
	'getMapMiFromHBase')
		URL="/v1/commonService/getMapMiFromHBase"
		export TABLE_NAME="inverter"
		JSON="$(jq -r '. += {"tableName": env.TABLE_NAME, "ps_key": env.PS_KEY, "beginTime": env.START_TIME, "endTime": env.END_TIME}' default.json)"
		;;

	# Invalid data
	'getValFromHBase')
		URL="/v1/commonService/getValFromHBase"
		export TABLE_NAME="inverter"
		export ROWKEY="inverter"
		JSON="$(jq -r '. += {"table": env.TABLE_NAME, "rowkey": env.ROWKEY}' default.json)"
		;;

	# Error
	'getStationInfoSql')
		URL="/v1/devService/getStationInfoSql"
		;;

	# Uses ps_key from
	'devicePointsDataFromMySql')
		URL="/v1/devService/devicePointsDataFromMySql"
		JSON="$(jq -r '. += {"ps_key": env.PS_KEY}' default.json)"
		;;

	# OK
	'getPListinfoFromMysql')
		URL="/v1/powerStationService/getPListinfoFromMysql"
		JSON="$(jq -r '. += {"psIds": env.PSID}' default.json)"
		;;

	# Error
	'getPowerStationTableDataSql')
		URL="/v1/devService/getPowerStationTableDataSql"
		;;

	# Error
	'getPowerStationTableDataSqlCount')
		URL="/v1/devService/getPowerStationTableDataSqlCount"
		;;

	# OK
	'getPListinfoFromMysql')
		URL="/v1/powerStationService/getPListinfoFromMysql"
		;;

	# Unknown args.
	'getTableDataSql')
		URL="/v1/devService/getTableDataSql"
		;;

	# Unknown args.
	'getTableDataSqlCount')
		URL="/v1/devService/getTableDataSqlCount"
		;;


	################################################################################
	'getPowerDeviceSetTaskDetailList')
		URL="/v1/devService/getPowerDeviceSetTaskDetailList"
		export QUERY_TYPE="1"
		if [ "${TASK_ID}" == "" ]; then TASK_ID="1566045"; fi
		#JSON="$(jq -r '. += {"ps_key": env.PS_KEY, "data_point": env.DATA_POINT, "data_type": env.DATA_TYPE, "query_type": env.QUERY_TYPE, "start_time": env.START_TIME, "end_time": env.END_TIME}' default.json)"
		JSON="$(jq -r '. += {"query_type": env.QUERY_TYPE, "task_id": env.TASK_ID}' default.json)"
		;;

	'getPowerDeviceSetTaskList')
		URL="/v1/devService/getPowerDeviceSetTaskList"
		export SIZE="256"
		export CURRENT_PAGE="$1"
		if [ "${CURRENT_PAGE}" == "" ]; then CURRENT_PAGE="1"; fi
		JSON="$(jq -r '. += {"size": env.SIZE, "curPage": env.CURRENT_PAGE}' default.json)"
		;;

	'getPsDataSupplementTaskList')
		URL="/v1/powerStationService/getPsDataSupplementTaskList"
		export SIZE="256"
		export CURRENT_PAGE="$1"
		if [ "${CURRENT_PAGE}" == "" ]; then CURRENT_PAGE="1"; fi
		JSON="$(jq -r '. += {"size": env.SIZE, "curPage": env.CURRENT_PAGE}' default.json)"
		;;

	'getRemoteUpgradeSubTasksList')
		URL="/v1/devService/getRemoteUpgradeSubTasksList"
		;;

	'getRemoteUpgradeTaskList')
		URL="/v1/devService/getRemoteUpgradeTaskList"
		;;

	'getModuleLogTaskList')
		URL="/integrationService/getModuleLogTaskList"
		#export SIZE="256"
		#export CURRENT_PAGE="$1"
		#if [ "${CURRENT_PAGE}" == "" ]; then CURRENT_PAGE="1"; fi
		# JSON="$(jq -r '. += {"curPage": env.CURRENT_PAGE}' default.json)"
		;;

	# Error
	'getSerialNum')
		URL="/v1/devService/getSerialNum"
		;;

	# Error
	'getPsCurveInfo')
		URL="/v1/devService/getPsCurveInfo"
		;;

	'getLoadCurveList')
		URL="/v1/reportService/getLoadCurveList"
		JSON="$(jq -r '. += {"ps_id": env.PSID, "monthDate": env.START_DATE}' default.json)"
		;;

	'getPsKpiForHoursByPsId')
		URL="/v1/powerStationService/getPsKpiForHoursByPsId"
		export DAY="1"
		JSON="$(jq -r '. += {"ps_id": env.PSID, "day": env.DAY}' default.json)"
		;;

	'getDevicePointMinuteDataList')
		URL="/v1/commonService/getDevicePointMinuteDataList"
		export POINTS="$1"
		if [ "${POINTS}" == "" ]; then POINTS="288"; fi
		JSON="$(jq -r '. += {"start_time_stamp": env.START_TIMESTAMP, "end_time_stamp": env.END_TIMESTAMP, "ps_key": env.PS_KEY, "points": env.POINTS}' default.json)"
		;;

	'getUpTimePoint')
		URL="/v1/devService/getUpTimePoint"
		;;

	'getDevicePoints')
		URL="/v1/devService/getDevicePoints"
		export POINT_ID="$1"
		if [ "${POINT_ID}" == "" ]; then POINT_ID="13003"; fi

		RESPONSEFILE="${API}-${POINT_ID}.response"
		JSON="$(jq -r '. += {"point_id": env.POINT_ID}' default.json)"
		;;

	'getApiCallsForAppkeys')
		URL="/v1/commonService/getApiCallsForAppkeys"
		;;

	'getAuthKey')
		URL="/v1/powerStationService/getAuthKey"
		;;

	'getPsAuthKey')
		URL="/v1/powerStationService/getPsAuthKey"
		JSON="$(jq -r '. += {"ps_id": env.PSID}' default.json)"
		;;

	# OK
	'getPowerStatistics')
		URL="/v1/powerStationService/getPowerStatistics"
		JSON="$(jq -r '. += {"ps_id": env.PSID}' default.json)"
		;;

	# OK
	'getPsList')
		URL="/v1/powerStationService/getPsList"
		;;

	# OK
	'createAppkeyInfo')
		URL="/v1/userService/createAppkeyInfo"
		export APP_KEY_NAME="MickMake"
		JSON="$(jq -r '. += {"appkey_name": env.APP_KEY_NAME, "user_account": env.USERID}' default.json)"
		;;

	# OK
	'exportPlantReportPDF')
		URL="/v1/powerStationService/exportPlantReportPDF"
		JSON="$(jq -r '. += {"ps_id": env.PSID}' default.json)"
		;;

	# OK
	'findPsType')
		URL="/v1/powerStationService/findPsType"
		JSON="$(jq -r '. += {"ps_id": env.PSID}' default.json)"
		;;

	# Error
	'findInfoByuuid')
		URL="/v1/devService/findInfoByuuid"
		export UUID="$1"
		if [ "${UUID}" == "" ]; then UUID="844775"; fi
		JSON="$(jq -r '. += {"uuid": env.UUID}' default.json)"
		;;

	# OK
	'findCurrentTask')
		URL="/v1/faultService/findCurrentTask"
		if [ "${TASK_ID}" == "" ]; then TASK_ID="1566045"; fi
		JSON="$(jq -r '. += {"taskId": env.TASK_ID}' default.json)"
		;;

	'executeTask')
		URL="/v1/faultService/executeTask"
		;;

	# OK
	'queryBatchCreatePsTaskList')
		URL="/v1/powerStationService/queryBatchCreatePsTaskList"
		;;

	# OK
	'queryAllPsIdAndName')
		URL="/v1/powerStationService/queryAllPsIdAndName"
		;;

	# Error
	'queryCtrlTaskById')
		URL="/v1/devService/queryCtrlTaskById"
		;;

	# UUID and TASK_ID from getPowerDeviceSetTaskList OR getPowerDeviceSetTaskDetailList
	'queryParamSettingTask')
		URL="/v1/devService/queryParamSettingTask"
		export QUERY_TYPE="1"
		export UUID="844775"
		if [ "${TASK_ID}" == "" ]; then TASK_ID="1566045"; fi
		JSON="$(jq -r '. += {"query_type": env.QUERY_TYPE, "task_id": env.TASK_ID, "uuid": env.UUID}' default.json)"
		;;

	# Requires task_id from getModuleLogTaskList
	'exportParamSettingValPDF')
		if [ "${TASK_ID}" == "" ]; then TASK_ID="1566045"; fi
		URL="/v1/devService/exportParamSettingValPDF"
		JSON="$(jq -r '. += {"task_id": env.TASK_ID}' default.json)"
		;;

	# OK
	'energyTrend')
		URL="/v1/powerStationService/energyTrend"
		;;

	# OK
	'communicationModuleDetail')
		URL="/v1/devService/communicationModuleDetail"
		JSON="$(jq -r '. += {"sn": env.SERIAL_NO}' default.json)"
		;;

	# Error
	'checkUnitStatus')
		URL="/v1/devService/checkUnitStatus"
		;;

	'saveEnvironmentCurve')
		URL="/v1/devService/saveEnvironmentCurve"
		;;


	################################################################################
	'updateTemplate')
		URL="/v1/devService/updateDataCurveTemplate"
		;;

	'getTemplateList')
		URL="/v1/devService/getTemplateList"
		;;

	'getTemplateByInfoType')
		URL="/v1/messageService/getTemplateByInfoType"
		export INFO_TYPE="1"	# Power Station Income Monthly Report-Mail Template
		export INFO_TYPE="2"	# Power Plant Revenue Annual Report-Mail Template
		export INFO_TYPE="3"	# Power station revenue report - email template
		export INFO_TYPE="4"	# Household Energy Storage Monthly Report-Mail Template
		export INFO_TYPE="5"	# Household Energy Storage Annual Report-Mail Template
		export INFO_TYPE="6"	# Household energy storage general report-mail template
		JSON="$(jq -r '. += {"info_type": env.INFO_TYPE}' default.json)"
		;;

	# OK - template_id from getTemplateList
	'queryUserCurveTemplateData')
		URL="/v1/devService/queryUserCurveTemplateData"
		export TEMPLATE_ID="7981"
		# JSON="$(jq -r '. += {"template_id": env.TEMPLATE_ID, "start_time": env.START_TIMESTAMP, "end_time": env.END_TIMESTAMP}' default.json)"

		export SELECT_DATE="2022-02-08 00:00~2022-02-08 23:59"
		export START_DATE="2022-02-08"
		export START_DATE="2022-02-08 00:00"
		export START_TIME="00:00"
		export START_TIME="20220208000000"

		export STOP_DATE="2022-02-08"
		export STOP_DATE="2022-02-08 23:59"
		export STOP_TIME="23:59"
		export STOP_TIME="20220208235959"

		JSON="$(jq -r '. += {"template_id": env.TEMPLATE_ID, "view_select_date": env.SELECT_DATE}' default.json)"

		JSON="$(echo "${JSON}" | jq -r '. += {"start_time_stamp": env.STARTT, "start_time": env.START_TIME, "begin_time": env.START_TIME, "startTime": env.START_TIME}')"
		JSON="$(echo "${JSON}" | jq -r '. += {"end_time_stamp":   env.STOPT,  "end_time":   env.STOP_TIME,  "endTime":   env.STOP_TIME, "endtime":   env.STOP_TIME}')"

		JSON="$(echo "${JSON}" | jq -r '. += {"start_date": env.START_DATE, "startDate": env.START_DATE, "startdate": env.START_DATE}')"
		JSON="$(echo "${JSON}" | jq -r '. += {"end_date":   env.STOP_DATE,  "endDate":   env.STOP_DATE, "enddate":   env.STOP_DATE, "date_type": "2"}')"
		;;

	################################################################################

	# Possible hourly data.
	'psHourPointsValue')
		URL="/v1/powerStationService/psHourPointsValue"
		JSON="$(jq -r '. += {"ps_id": env.PSID}' default.json)"
		;;

	'powerDevicePointList')
		URL="/v1/reportService/powerDevicePointList"
		;;

	'queryDevicePointMinuteDataList')
		URL="/v1/commonService/queryDevicePointMinuteDataList"
		export POINTS="$1"
		if [ "${POINTS}" == "" ]; then POINTS="13150"; fi
		JSON="$(jq -r '. += {"start_time_stamp": env.START_TIMESTAMP, "end_time_stamp": env.END_TIMESTAMP, "ps_key": env.PS_KEY, "points": env.POINTS}' default.json)"
		;;

	'queryDevicePointsDayMonthYearDataList')
		URL="/v1/commonService/queryDevicePointsDayMonthYearDataList"
		export DATA_POINT="81012"
		export DATA_TYPE="17"
		export QUERY_TYPE="1"
		JSON="$(jq -r '. += {"ps_key": env.PS_KEY, "data_point": env.DATA_POINT, "data_type": env.DATA_TYPE, "query_type": env.QUERY_TYPE, "start_time": env.START_TIME, "end_time": env.END_TIME}' default.json)"
		;;

	'queryUnitList')
		URL="/v1/userService/queryUnitList"
		;;

	'showPSView')
		URL="/v1/powerStationService/showPSView"
		JSON="$(jq -r '. += {"ps_id": env.PSID}' default.json)"
		;;

	*)
		echo "Unknown API endpoint..."
		grep "')$" get.sh
		exit
		;;
esac

echo "${JSON}" | jq -r > ${POSTFILE}

curl -X POST -H 'Content-Type: application/json' -A "${USER_AGENT}" --stderr error.log \
	--data-binary "@${POSTFILE}" \
	https://augateway.isolarcloud.com${URL} | tee ${RESPONSEFILE}

