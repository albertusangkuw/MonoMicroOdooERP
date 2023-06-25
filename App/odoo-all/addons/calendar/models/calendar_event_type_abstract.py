
from abc import ABC, abstractmethod
import logging
import requests
from odoo import http, models
import json
_logger = logging.getLogger(__name__)
from datetime import datetime

class MeetingTypeAbstract(models.AbstractModel):
    _name = 'calendar.event.type_abstract'
    _description = ' Abstract RPC'
    _TABLE_NAME = "calendar_event_type"
    _MODEL_NAME = "calendar.event.type"
    _HOST_RPC = 'http://172.18.0.52:1324'
    _PATH_RPC = "/jsonrpc/web/dataset/call_kw/calendar.event.type/"
    _TARGET_COLUMN = [ 'id', 'color',  'create_uid', 'write_uid', 'name' ]
    _TARGET_TYPE_COLUMN = [ int , int , int , int, str]    
    
    def create_payload(self,modelName, method,args, context, idCall):
        return {
            "id": idCall,
            "jsonrpc": "2.0",
            "method": "call",
            "params": {
                "model": modelName,
                "method": method,
                "args": args,
                "kwargs": {
                    "context": context,
                },
            },
        }
    
    def _is_same_data(self, old, new):
        for i,c in enumerate(self._TARGET_COLUMN):
           if old[c] != self.convert_to_odoo_type(new[c],self._TARGET_TYPE_COLUMN[i]):
               return  False
        return True
      
    def call_rpc(self, method, headers, context , args=[],keyResult='result'):
        payload = self.create_payload(self._MODEL_NAME, method,args,context,1)
        # print(json.dumps(payload))
        try:
            _logger.info(f"Request RPC: {self._HOST_RPC}{self._PATH_RPC}")
            response = requests.post(f"{self._HOST_RPC}{self._PATH_RPC}", headers=headers, data=json.dumps(payload))
            if response.ok:
                result = response.json()
                if 'error' in result:
                    raise Exception(f"Internal Error In Result RPC : {result['error']}")
                # print(result)
                return result[keyResult]
            else:
                raise Exception(f"Internal Error From RPC : {self._HOST_RPC}{self._PATH_RPC} - Response Not OK")
        except:
            raise Exception(f"Internal Error By Request RPC: {self._HOST_RPC}{self._PATH_RPC}  ")
  
    def convert_to_odoo_type(self,val,targetVal):
        if isinstance(val, list):
            # [ 1, "My Label"]
            if type(val[0]) == targetVal:
                return val[0]
            else :
                print("Unknown type: ", val, type(val))
                return val[0]
        else: 
            if type(val) == targetVal:
                return val
            elif type(val) == bool:
                # "id": False , => "id": [1, "My Label"]
                return None
            elif type(val) == str and targetVal == "datetime":
                # "date_str": 2023-05-19 14:52:07 , => "timestamp": tmpstamp(....)
                time_str = val
                format_str = "%Y-%m-%d %H:%M:%S"
                timestamp = datetime.strptime(time_str, format_str)
                return timestamp
            else:
                print("Unknown type:", type(val))
                return val
     
    def manual_insert(self,env,records):
        if len(records) == 0:
            return
        try:
            listColumn = ",".join(self._TARGET_COLUMN)
            query = f"INSERT INTO {self._TABLE_NAME} ({listColumn}) VALUES "
            targetValues =  f"({', '.join(['%s'] * len(self._TARGET_COLUMN))})"
            
            values = []
            realValues = []
            for record in records:
                values.append(targetValues)
                tmpRV = []
                c = self._TARGET_COLUMN
                ct = self._TARGET_TYPE_COLUMN
                for i in range(0,len(self._TARGET_COLUMN)):
                    newData = record.get(c[i])
                    tmpRV.append(self.convert_to_odoo_type(newData, ct[i]))
                realValues += tmpRV
            query += ','.join(values)
            
            placeholders = tuple(realValues)
            preview_sql_query = query % placeholders
            print(preview_sql_query)
            env.cr.execute(query, realValues)
            env.cr.commit()
            _logger.info('Data inserted successfully into %s', self._TABLE_NAME)
        except Exception as e:
            _logger.error('Error inserting data into %s: %s', self._TABLE_NAME, str(e))
            raise
    
    def manual_update(self,env,records):
        if len(records) == 0:
            return
        try:
            clTmp = []
            for c in self._TARGET_COLUMN:
                # ID tidak bisa diupdate
                if c == "id":
                    continue
                clTmp.append("" + c + "= %s" )
            targetValues = ", ".join(clTmp)
            query = f"UPDATE {self._TABLE_NAME} SET {targetValues} "
            query += " WHERE id = %s ;" 
            
            for record in records:
                realValues = []
                
                c = self._TARGET_COLUMN
                ct = self._TARGET_TYPE_COLUMN
                targetID = 0 
                for i in range(0,len(self._TARGET_COLUMN)):
                    # ID tidak bisa diupdate
                    if c[i] == "id":
                        targetID = record.get(c[i])
                        continue
                    newData = record.get(c[i])
                    realValues.append(self.convert_to_odoo_type(newData,ct[i]))
                realValues.append(targetID)
                
                placeholders = tuple(realValues)
                # preview_sql_query = query % placeholders
                # print(preview_sql_query)
                env.cr.execute(query,realValues)
                
            env.cr.commit()
            _logger.info('Data updated successfully in %s', self._TABLE_NAME)
        except Exception as e:
            _logger.error('Error updating data in %s: %s', self._TABLE_NAME, str(e))
            raise
    
    def manual_get(self,env,ids=[],returnobj=False):
        try:
            columns = ', '.join(self._TARGET_COLUMN)
            query = "SELECT %s FROM %s" % (columns, self._TABLE_NAME)
            if len(ids) > 0:
                idsStr =  ",".join(str(id) for id in ids)
                query += f" WHERE id IN ({ idsStr })";
            
            env.cr.execute(query)
            results = env.cr.fetchall()
            if returnobj:
                records = {}
                for row in results:
                    record = {}
                    for i, column in enumerate(self._TARGET_COLUMN):
                        record[column] = row[i]
                    # Asumsi ID selalu pertama
                    records[row[0]] = record
                return records
            else:
                records = []
                for row in results:
                    record = {}
                    for i, column in enumerate(self._TARGET_COLUMN):
                        record[column] = row[i]
                    records.append(record)

                return records
        except Exception as e:
            _logger.error('Error selecting records from %s: %s', self._TABLE_NAME, str(e))
            raise
    
    def manual_count(self,env):
        cr = env.cr
        query = f"SELECT COUNT(*) FROM {self._TABLE_NAME}"
        cr.execute(query)
        result = cr.fetchone()
        count = 0
        if result:
            count = result[0]
        return count
     
    def manual_delete(self,env,records):
        cr = env.cr
        for id in records:
            cr.execute('DELETE FROM ' + self._TABLE_NAME +' WHERE id = %s', (id,))
        cr.commit()
        return
    
    @abstractmethod
    def sync_data(self,ids):
        pass