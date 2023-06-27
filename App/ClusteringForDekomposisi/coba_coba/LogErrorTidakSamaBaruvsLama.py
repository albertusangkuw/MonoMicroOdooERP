--- fungsi original start
def searchParentCall(arrStr):
    tmp = ""
    all = arrStr.copy()
    #Edge Case 1
    if len(arrStr) > 1 and  f'{arrStr[0]}.{arrStr[1]}' == 'odoo.Command':
        arrStr[1] = 'cli'
        arrStr.insert(2, 'command')

    while(len(arrStr) > 0):
        arrStr.pop()
        parent = ".".join(arrStr)
        if parent in callGraphWeight:
            return parent
    
    #Edge Case 2
    arrStr = all.copy()
    while(len(arrStr) > 0):
        arrStr.pop()
        parent = f'{".".join(arrStr)}.__init__'
        if parent in callGraphWeight:
            return parent
    # 'odoo', 'addons', 'test_main_flows', '_auto_install_enterprise_dependencies']    
    # print("Parent Not Found : " , all )
    return tmp
def updateCGWeight(callGraphFiltered,callGraphWeight):
    #Versi Per Addons
    def checkIsBP(m):
        blackLS = ['test', 'l10n']
        for i in blackLS:
            if i in m:
                return True
        return False
    for c,r in callGraphFiltered.items():
        arrC = c.split(".")
        newC = c
        if checkIsBP(c):
            continue
        if arrC[0] != "addons":
            continue
        
        if c not in callGraphWeight:
            newC = searchParentCall(arrC)
        
        for rc , v in r.items():
            if checkIsBP(rc):
                continue
            newRC = rc
            if rc not in callGraphWeight:
                newRC = searchParentCall(rc.split("."))
            if newRC == "" or newC == "":
                continue   
            if newRC not in callGraphWeight[newC]:
                callGraphWeight[newC][newRC] = 0
            callGraphWeight[newC][newRC] += v
    return callGraphWeight
--- fungsi original end
def funsinnya():
    for x in callGraphWeight:
        if x not in newCG:
            if len(odooTree.countCall(x)) == 0:
                continue
            else:
                print("Key not found in  " , x)
        elif callGraphWeight[x] != newCG[x]:
            # if callGraphWeight[x][x] != newCG[x][x]:
            #     print("data tidak sama untuk sendiri")
            #     continue
            print("Data not same in new " , x  )
            print(callGraphWeight[x])
            print("-vs-")
            print(newCG[x])
            print("-------end ")

Summarnya :
Key not found in   addons.base
Data not same in new  addons.base.wizard.base_partner_merge
Data not same in new  addons.base.models.res_lang
Data not same in new  addons.base.models.ir_model
Data not same in new  addons.base.models.res_country
Data not same in new  addons.base.models.ir_demo_failure
Data not same in new  addons.base.models.res_bank
Data not same in new  addons.base.models.ir_filters
Data not same in new  addons.base.models.ir_sequence
Data not same in new  addons.base.models.ir_exports
Data not same in new  addons.base.models.ir_ui_menu
Data not same in new  addons.base.models.ir_ui_view
Data not same in new  addons.base.models.res_company
Data not same in new  addons.base.models.res_currency
Data not same in new  addons.base.models.ir_module
Data not same in new  addons.base.models.res_users
Data not same in new  addons.base.models.ir_cron
Data not same in new  addons.base.models.ir_actions
    

hr_work_entry_contract ??? apa ini kenapa banyak bikin relasi tidak nyambung ?
Key not found in   addons.attachment_indexation
Key not found in   addons.board
Key not found in   addons.hw_drivers
Key not found in   addons.hw_escpos
Key not found in   addons.website_sms
Key not found in   addons.google_account
Key not found in   addons.base
Data not same in new  addons.base.wizard.base_partner_merge
{'addons.base.wizard.base_partner_merge': 27, 'addons.base.models.res_users': 2}
-vs-
{'addons.base.models.res_users': 2, 'addons.base.wizard.base_partner_merge': 26}
-------end 
Data not same in new  addons.base.models.res_lang
{'addons.base.models.res_lang': 15, 'addons.hr_work_entry_contract': 1, 'addons.base.models.res_users': 2}
-vs-
{'addons.base.models.res_users': 2, 'addons.base.models.res_lang': 15}
-------end 
Key not found in   addons.base.models.ir_mail_server
Data not same in new  addons.base.models.ir_model
{'addons.base.models.ir_model': 75, 'addons.hr_work_entry_contract': 9, 'addons.base.models.res_users': 2, 'addons.base.models.ir_ui_menu': 1}
-vs-
{'addons.base.models.res_users': 2, 'addons.base.models.ir_ui_menu': 1, 'addons.base.models.ir_model': 75}
-------end 
Key not found in   addons.base.models.ir_autovacuum
Key not found in   addons.base.models.ir_http
Data not same in new  addons.base.models.res_country
{'addons.base.models.res_country': 1, 'addons.base.models.res_users': 2}
-vs-
{'addons.base.models.res_users': 2}
-------end 
Data not same in new  addons.base.models.ir_demo_failure
{'addons.base.models.res_users': 2, 'addons.base.models.ir_demo_failure': 1}
-vs-
{'addons.base.models.res_users': 2}
-------end 
Data not same in new  addons.base.models.res_bank
{'addons.base.models.res_bank': 5, 'addons.base.models.res_users': 2, 'addons.base.models.res_currency': 1, 'addons.base.models.res_partner': 1}
-vs-
{'addons.base.models.res_bank': 4, 'addons.base.models.res_users': 2, 'addons.base.models.res_currency': 1, 'addons.base.models.res_partner': 1}
-------end 
Key not found in   addons.base.models.ir_binary
Data not same in new  addons.base.models.ir_filters
{'addons.base.models.ir_filters': 5, 'addons.hr_work_entry_contract': 1, 'addons.base.models.ir_actions': 1, 'addons.base.models.res_users': 3}
-vs-
{'addons.base.models.ir_actions': 1, 'addons.base.models.res_users': 3, 'addons.base.models.ir_filters': 5}
-------end 
Data not same in new  addons.base.models.ir_sequence
{'addons.base.models.ir_sequence': 31, 'addons.hr_work_entry_contract': 1, 'addons.base.models.res_users': 2}
-vs-
{'addons.base.models.res_users': 2, 'addons.base.models.ir_sequence': 30}
-------end 
Data not same in new  addons.base.models.ir_exports
{'addons.base.models.res_users': 2, 'addons.base.models.ir_exports': 1}
-vs-
{'addons.base.models.res_users': 2}
-------end 
Data not same in new  addons.base.models.ir_ui_menu
{'addons.base.models.ir_ui_menu': 12, 'addons.hr_work_entry_contract': 2, 'addons.base.models.res_users': 3}
-vs-
{'addons.base.models.ir_ui_menu': 10, 'addons.base.models.res_users': 3}
-------end 
Key not found in   addons.base.models.ir_qweb_fields
Data not same in new  addons.base.models.ir_ui_view
{'addons.base.models.res_users': 3, 'addons.base.models.ir_ui_view': 128, 'addons.hr_work_entry_contract': 9}
-vs-
{'addons.base.models.res_users': 3, 'addons.base.models.ir_ui_view': 127, 'addons.hr_work_entry_contract': 9}
-------end 
Data not same in new  addons.base.models.res_company
{'addons.base.models.res_company': 6, 'addons.base.models.res_country': 2, 'addons.base.models.res_users': 3, 'addons.base.models.res_currency': 1, 'addons.base.models.ir_ui_view': 1, 'addons.base.models.report_paperformat': 1, 'addons.base.models.res_partner': 1}
-vs-
{'addons.base.models.res_company': 4, 'addons.base.models.res_country': 2, 'addons.base.models.res_users': 3, 'addons.base.models.res_currency': 1, 'addons.base.models.ir_ui_view': 1, 'addons.base.models.report_paperformat': 1, 'addons.base.models.res_partner': 1}
-------end 
Data not same in new  addons.base.models.res_currency
{'addons.base.models.res_currency': 14, 'addons.base.models.res_company': 1, 'addons.base.models.res_users': 2}
-vs-
{'addons.base.models.res_company': 1, 'addons.base.models.res_users': 2, 'addons.base.models.res_currency': 16}
-------end 
Key not found in   addons.base.models.ir_qweb
Data not same in new  addons.base.models.ir_module
{'addons.base.models.ir_module': 28, 'addons.hr_work_entry_contract': 5, 'addons.base.models.res_users': 2}
-vs-
{'addons.base.models.res_users': 2, 'addons.base.models.ir_module': 27}
-------end 
Data not same in new  addons.base.models.res_users
{'addons.base.models.res_users': 82, 'addons.hr_work_entry_contract': 2, 'addons.base.models.res_partner': 2, 'addons.base.models.ir_actions': 1, 'addons.base.models.res_company': 2}
-vs-
{'addons.base.models.res_partner': 2, 'addons.base.models.ir_actions': 1, 'addons.base.models.res_company': 2, 'addons.base.models.res_users': 78}
-------end 
Data not same in new  addons.base.models.ir_cron
{'addons.base.models.res_users': 2, 'addons.base.models.ir_cron': 10}
-vs-
{'addons.base.models.res_users': 2, 'addons.base.models.ir_cron': 9}
-------end 
Key not found in   addons.base.models.ir_profile
Key not found in   addons.base.models.assetsbundle
Data not same in new  addons.base.models.ir_actions
{'addons.base.models.ir_actions': 9, 'addons.base.models.ir_model': 1, 'addons.base.models.res_users': 2}
-vs-
{'addons.base.models.ir_model': 1, 'addons.base.models.res_users': 2, 'addons.base.models.ir_actions': 8}
-------end 
Key not found in   addons.base.models.decimal_precision
Key not found in   addons.base.controllers.rpc

