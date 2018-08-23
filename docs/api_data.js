define({ "api": [  {    "type": "delete",    "url": "/shop/rules",    "title": "Delete one to many search rules by id",    "name": "DeleteRules",    "group": "Rules",    "parameter": {      "fields": {        "Parameter": [          {            "group": "Parameter",            "type": "String",            "optional": false,            "field": "id",            "description": "<p>The id of the rule you want to delete. You can pass as many id params as desired.</p>"          },          {            "group": "Parameter",            "type": "String",            "optional": false,            "field": "shopify",            "description": "<p>The shopify store i.e. fastseer.myshopify.com</p>"          },          {            "group": "Parameter",            "type": "String",            "optional": false,            "field": "hmac",            "description": "<p>Secure hmac param from Shopify admin.</p>"          },          {            "group": "Parameter",            "type": "String",            "optional": false,            "field": "locale",            "description": "<p>Locale from the Shopify admin.</p>"          },          {            "group": "Parameter",            "type": "Number",            "optional": false,            "field": "timestamp",            "description": "<p>Timestamp from the Shopify admin.</p>"          }        ]      }    },    "success": {      "fields": {        "Success 200": [          {            "group": "Success 200",            "type": "Object",            "optional": false,            "field": "MessageResponse",            "description": "<p>See <code>MessageResponse</code> type.</p>"          }        ]      }    },    "error": {      "fields": {        "Error 4xx": [          {            "group": "Error 4xx",            "type": "Object",            "optional": false,            "field": "ErrorResponse",            "description": "<p>See <code>ErrorResponse</code> type.</p>"          }        ]      }    },    "version": "0.0.0",    "filename": "./api_admin_rules.go",    "groupTitle": "Rules"  },  {    "type": "get",    "url": "/shop/rules",    "title": "Search and retrieve rules",    "name": "GetRules",    "group": "Rules",    "parameter": {      "fields": {        "Parameter": [          {            "group": "Parameter",            "type": "String",            "optional": false,            "field": "q",            "description": "<p>The query to search for rules</p>"          },          {            "group": "Parameter",            "type": "String",            "optional": false,            "field": "shopify",            "description": "<p>The shopify store i.e. fastseer.myshopify.com</p>"          },          {            "group": "Parameter",            "type": "String",            "optional": false,            "field": "hmac",            "description": "<p>Secure hmac param from Shopify admin.</p>"          },          {            "group": "Parameter",            "type": "String",            "optional": false,            "field": "locale",            "description": "<p>Locale from the Shopify admin.</p>"          },          {            "group": "Parameter",            "type": "Number",            "optional": false,            "field": "timestamp",            "description": "<p>Timestamp from the Shopify admin.</p>"          }        ]      }    },    "success": {      "fields": {        "Success 200": [          {            "group": "Success 200",            "type": "Object",            "optional": false,            "field": "SearchRules",            "description": "<p>Returns a json array of <code>SearchRule</code>.</p>"          }        ]      }    },    "error": {      "fields": {        "Error 4xx": [          {            "group": "Error 4xx",            "type": "Object",            "optional": false,            "field": "ErrorResponse",            "description": "<p>See <code>ErrorResponse</code> type.</p>"          }        ]      }    },    "version": "0.0.0",    "filename": "./api_admin_rules.go",    "groupTitle": "Rules"  },  {    "type": "put",    "url": "/shop/rules",    "title": "Put one to many search rules",    "name": "PutRules",    "group": "Rules",    "parameter": {      "fields": {        "Parameter": [          {            "group": "Parameter",            "type": "Object",            "optional": false,            "field": "rules",            "description": "<p>The request body, accepts a json <code>[]rules.SearchRule</code>. Leave id blank if you are creating a new rule. Otherwise, and existing rule will be overwritten if the ids match.</p>"          },          {            "group": "Parameter",            "type": "String",            "optional": false,            "field": "shopify",            "description": "<p>The shopify store i.e. fastseer.myshopify.com</p>"          },          {            "group": "Parameter",            "type": "String",            "optional": false,            "field": "hmac",            "description": "<p>Secure hmac param from Shopify admin.</p>"          },          {            "group": "Parameter",            "type": "String",            "optional": false,            "field": "locale",            "description": "<p>Locale from the Shopify admin.</p>"          },          {            "group": "Parameter",            "type": "Number",            "optional": false,            "field": "timestamp",            "description": "<p>Timestamp from the Shopify admin.</p>"          }        ]      },      "examples": [        {          "title": "Request Body Example:",          "content": "\t\t[\n      {\n        \"id\":\"b1f3cd69-9d44-404e-a97b-993568391988\",\n        \"name_s\":\"ipad sale\",\n        \"actAddBqs_ss\":[\"id:1234^10\"],\n        \"order_i\":1},\n      {\n        \"id\":\"dd75c626-c62b-4595-bae5-9692f9449408\",\n        \"name_s\":\"iphone promotion\",\n        \"order_i\":1}]",          "type": "json"        }      ]    },    "success": {      "fields": {        "Success 200": [          {            "group": "Success 200",            "type": "Object",            "optional": false,            "field": "MessageResponse",            "description": "<p>See <code>MessageResponse</code> type.</p>"          }        ]      }    },    "error": {      "fields": {        "Error 4xx": [          {            "group": "Error 4xx",            "type": "Object",            "optional": false,            "field": "ErrorResponse",            "description": "<p>See <code>ErrorResponse</code> type.</p>"          }        ]      }    },    "version": "0.0.0",    "filename": "./api_admin_rules.go",    "groupTitle": "Rules"  },  {    "type": "post",    "url": "/shop/theme/install",    "title": "Reinstall theme assets into the current shopify theme",    "name": "PostThemeAssets",    "group": "Theme",    "parameter": {      "fields": {        "Parameter": [          {            "group": "Parameter",            "type": "String",            "optional": false,            "field": "shopify",            "description": "<p>The shopify store i.e. fastseer.myshopify.com</p>"          },          {            "group": "Parameter",            "type": "String",            "optional": false,            "field": "hmac",            "description": "<p>Secure hmac param from Shopify admin.</p>"          },          {            "group": "Parameter",            "type": "String",            "optional": false,            "field": "locale",            "description": "<p>Locale from the Shopify admin.</p>"          },          {            "group": "Parameter",            "type": "Number",            "optional": false,            "field": "timestamp",            "description": "<p>Timestamp from the Shopify admin.</p>"          }        ]      }    },    "success": {      "fields": {        "Success 200": [          {            "group": "Success 200",            "type": "Object",            "optional": false,            "field": "MessageResponse",            "description": "<p>See <code>MessageResponse</code> type.</p>"          }        ]      }    },    "error": {      "fields": {        "Error 4xx": [          {            "group": "Error 4xx",            "type": "Object",            "optional": false,            "field": "ErrorResponse",            "description": "<p>See <code>ErrorResponse</code> type.</p>"          }        ]      }    },    "version": "0.0.0",    "filename": "./api_admin_reinstallthemeassets.go",    "groupTitle": "Theme"  },  {    "success": {      "fields": {        "Success 200": [          {            "group": "Success 200",            "optional": false,            "field": "varname1",            "description": "<p>No type.</p>"          },          {            "group": "Success 200",            "type": "String",            "optional": false,            "field": "varname2",            "description": "<p>With type.</p>"          }        ]      }    },    "type": "",    "url": "",    "version": "0.0.0",    "filename": "./docs/main.js",    "group": "_Users_evanpease_Development_go_src_github_com_ezeev_fastseer_docs_main_js",    "groupTitle": "_Users_evanpease_Development_go_src_github_com_ezeev_fastseer_docs_main_js",    "name": ""  }] });
