# -*- coding: utf-8 -*-
# Part of Odoo. See LICENSE file for full copyright and licensing details.
from odoo import api, fields, models, _
from odoo.exceptions import ValidationError, UserError
from odoo import http
import requests
import logging
import json
from datetime import datetime
from .pos_category_abstract import PosCategoryAbstract
_logger = logging.getLogger(__name__)


class PosCategoryAdapter(models.Model, PosCategoryAbstract):
    _name = "pos.category"
    _description = "Point of Sale Category"
    _order = "sequence, name"
    _rec_name = "display_name"

    name = fields.Char(string='Category Name')
    display_name = fields.Char(string='Display Name')
    parent_id = fields.Many2one('pos.category', string='Parent Category')
    child_id = fields.One2many('pos.category', 'parent_id', string='Children Categories')
    sequence = fields.Integer(help="Gives the sequence order when displaying a list of product categories.")
    image_128 = fields.Image("Image", max_width=128, max_height=128)
    has_image = fields.Boolean()

     
    def sync_data(self,ids = []):
        _logger.info("Sync Data Started !!" )
        # Kondisi ketika ID's yang diminta semua maka ID request dihilangkan
        if len(ids) == self.manual_count(self.env):
            ids = []
        api_data =  self.call_rpc("read", dict(http.request.httprequest.headers), self.env.context ,[ids])
        existing_records = self.manual_get(self.env,ids=ids,returnobj=True)
        createList = []
        updateList = []
        sameCounter = 0
        
        for data in api_data:
            record = {}
            if data['id'] in existing_records:
                existing_records[data['id']]['___metadata_found'] = True
                record = existing_records[data['id']]
            if record:
                if self._is_same_data(record,data) == False:
                    updateList.append(data)
                else:
                    sameCounter += 1
            else:
                createList.append(data)
            
        deleteList = []
        
        # #Delete Cek, ketika semua data diambil dari service
        ex_key = list(existing_records.keys())
        for k in ex_key:
            if '___metadata_found' in existing_records[k]:
                #Sudah ditemukan sebelumnya ketika pencarian data di api_data
                continue
            else: 
                deleteList.append(k)
        
        _logger.info(f"Existing: {len(existing_records)}; Create: {len(createList)}; Update: {len(updateList)} , Delete: {len(deleteList)}, Unchange List: {sameCounter} ")
        
        self.manual_delete(self.env,deleteList)
        self.manual_insert(self.env,createList)
        self.manual_update(self.env,updateList)
        self.env[self._name].clear_caches()
   
    def _read(self ,field_names):
        print("PRINT: def _read internal called:" , field_names  , self._table  , self.ids)
        # Update data keseluruhan, diperlukan mekanisme filter ID 
        self.sync_data(self.ids)
        return super(PosCategoryAdapter,self)._read(field_names)
    
    @api.model
    def create(self, records):
        _logger.info("Create Triggered from Mono- ")
        newID =  self.call_rpc("create", dict(http.request.httprequest.headers), self.env.context ,[records])
        return self.browse(newID)
        
    def write(self, vals):
        if not self:
            return True
        _logger.info("Write Triggered from Mono- ")
        self.call_rpc("write", dict(http.request.httprequest.headers),  self.env.context,[self.ids,vals])
        return True
        
    def unlink(self):
        if not self:
            return True
        _logger.info("Unlink: ")
        self.call_rpc("unlink", dict(http.request.httprequest.headers), self.env.context ,[self.ids])
        return True
    
    @api.model
    def browse(self, ids=None):
        if not ids:
            ids = ()
        elif ids.__class__ is int:
            ids = (ids,)
        else:
            ids = tuple(ids)
        _logger.info("Browse Adapter Triggered :" + str(ids))
        currRow  = self.manual_count(self.env)
        if currRow == 0:
            self.sync_data()  
        return self.__class__(self.env, ids, ids)
    
    @api.model
    def read(self,fields=None, load='_classic_read'):
        _logger.info("Read is called  ")
        for name in fields:
            field = self._fields.get(name)
            if not field:
                # Jika ada field yang tidak disimpan di db make panggil api, contohnya computer field 
                return self.call_rpc("read", dict(http.request.httprequest.headers), self.env.context,
                                    [fields])
        return super(PosCategoryAdapter,self).read(fields, load)
    

        