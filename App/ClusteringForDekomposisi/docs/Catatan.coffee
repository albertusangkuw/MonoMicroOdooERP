Module: note
---------------
{'__module__': 'addons.note.models.note', 
'_name': 'note.note', 
'_inherit': ['mail.thread', 'mail.activity.mixin'], 
'_description': 'Note', '_order': 'sequence, id desc', 
'_get_default_stage_id': <function Note._get_default_stage_id at 0x7f6a8e38fd00>, 
'name': note.note.name, 
'company_id': note.note.company_id, 
'user_id': note.note.user_id, 
'memo': note.note.memo, 
'sequence': note.note.sequence, 
'stage_id': note.note.stage_id, 
'stage_ids': note.note.stage_ids, 
'open': note.note.open, 
'date_done': note.note.date_done, 
'color': note.note.color, 
'tag_ids': note.note.tag_ids, 
'message_partner_ids': note.note.message_partner_ids, 
'_compute_name': <function Note._compute_name at 0x7f6a8e3e09d0>, 
'_compute_stage_id': <function Note._compute_stage_id at 0x7f6a8e3e0940>, 
'_inverse_stage_id': <function Note._inverse_stage_id at 0x7f6a8e3e0a60>, 
'name_create': <function Note.name_create at 0x7f6a8e3e0af0>, 
'read_group': <function Note.read_group at 0x7f6a8e3e0b80>, 
'action_close': <function Note.action_close at 0x7f6a8e3e0c10>, 
'action_open': <function Note.action_open at 0x7f6a8e3e0ca0>, 
'write': <function Note.write at 0x7f6a8e3e0d30>, 
'__slots__': (),
 '_field_definitions': [note.note.name, note.note.company_id, note.note.user_id, note.note.memo, note.note.sequence, note.note.stage_id, note.note.stage_ids, note.note.open, note.note.date_done, note.note.color, note.note.tag_ids, note.note.message_partner_ids, note.note.id, note.note.__last_update, note.note.display_name, note.note.create_uid, note.note.create_date, note.note.write_uid, note.note.write_date], '_module': 'models', '__doc__': None, 'id': note.note.id, '__last_update': note.note.__last_update, 'display_name': note.note.display_name, 'create_uid': note.note.create_uid, 'create_date': note.note.create_date, 'write_uid': note.note.write_uid, 'write_date': note.note.write_date, '<lambda>': 1.0}
{'__module__': 'addons.note.models.note', '_name': 'note.stage', '_description': 'Note Stage', '_order': 'sequence', 'name': note.stage.name, 'sequence': note.stage.sequence, 'user_id': note.stage.user_id, 'fold': note.stage.fold, '__slots__': (), '_field_definitions': [note.stage.name, note.stage.sequence, note.stage.user_id, note.stage.fold, note.stage.id, note.stage.__last_update, note.stage.display_name, note.stage.create_uid, note.stage.create_date, note.stage.write_uid, note.stage.write_date], '_module': 'models', '__doc__': None, 'id': note.stage.id, '__last_update': note.stage.__last_update, 'display_name': note.stage.display_name, 'create_uid': note.stage.create_uid, 'create_date': note.stage.create_date, 'write_uid': note.stage.write_uid, 'write_date': note.stage.write_date, '<lambda>': 1.0}
{'__module__': 'addons.note.models.note', '_name': 'note.tag', '_description': 'Note Tag', 'name': note.tag.name, 'color': note.tag.color, '_sql_constraints': [('name_uniq', 'unique (name)', 'Tag name already exists !')], '__slots__': (), '_field_definitions': [note.tag.name, note.tag.color, note.tag.id, note.tag.__last_update, note.tag.display_name, note.tag.create_uid, note.tag.create_date, note.tag.write_uid, note.tag.write_date], '_module': 'models', '__doc__': None, 'id': note.tag.id, '__last_update': note.tag.__last_update, 'display_name': note.tag.display_name, 'create_uid': note.tag.create_uid, 'create_date': note.tag.create_date, 'write_uid': note.tag.write_uid, 'write_date': note.tag.write_date, '<lambda>': 1.0}
{'Note': {'name': 'note.note', 'inherit': ['mail.thread', 'mail.activity.mixin'], 'inherits': {}, 'attribute_rel': {'company_id': 'res.company', 'create_uid': 'res.users', 'stage_id': 'note.stage', 'stage_ids': 'note.stage', 'tag_ids': 'note.tag', 'user_id': 'res.users', 'write_uid': 'res.users'}}, 'Stage': {'name': 'note.stage', 'inherit': (), 'inherits': {}, 'attribute_rel': {'create_uid': 'res.users', 'user_id': 'res.users', 'write_uid': 'res.users'}}, 'Tag': {'name': 'note.tag', 'inherit': (), 'inherits': {}, 'attribute_rel': {'create_uid': 'res.users', 'write_uid': 'res.users'}}}


callGraphWeightSorted['addons.note']

sum
{'odoo.http': 1,
 'odoo.api': 4,
 'odoo.models': 6,
 'addons.note': 8,
 'odoo.__init__': 1,
 'odoo.fields': 16,
 'odoo.tools.__init__': 1,
 'odoo.addons.__init__': 1,
 'addons.mail': 4,
 'odoo.addons.base.models.res_company': 1,
 'odoo.addons.base.models.res_users': 9}


inspect
 {
 'addons.note': 8,
 'addons.mail': 4,
 'odoo.addons.base.models.res_company': 1,
 'odoo.addons.base.models.res_users': 9}


 expected
 {
    'addons.note': 3 ??,
    'addons.mail': 2,
    'odoo.addons.base.models.res_company': 1,
    'odoo.addons.base.models.res_users': 8
    'addons.web_editor': 1,
}

{'odoo.http': 2,
 'odoo.api': 8,
 'odoo.models': 12,
 'addons.note': 12,
 'odoo.__init__': 2,
 'odoo.fields': 32,
 'odoo.tools.__init__': 2,
 'addons.web_editor': 2,
 'addons.mail': 4,
 'odoo.addons.base.models.res_company': 1,
 'odoo.addons.base.models.res_users': 9}

print(listModuleName['note.note'])
print(listModuleName['note.tag'])
print(listModuleName['note.stage'])
print(listModuleName['addons.note.models.mail_activity'])
print(listModuleName['addons.note.models.mail_activity_type'])
print(listModuleName['addons.note.models.res_users'])
 
 addons.note.models.res_users.Users._init_data_user_note_stages  =>  {'addons.note.models.res_users.Users.env.cr.execute': 1,
                                                                 'addons.note.models.res_users.Users.env.cr.fetchall': 1}
 addons.note.models.note.Note.write  =>  {'addons.web_editor.controllers.main.handle_history_divergence': 1}
 addons.note.models.note.Note.action_close  =>  { 'addons.note.models.note.Note.write': 1}
 addons.note.models.note.Note.action_open  =>  {'addons.note.models.note.Note.write': 1}




 Self 8

 {'addons.web_editor': 1,
 'addons.mail': 4,
 'odoo.addons.base.models.res_company': 1,
 'odoo.addons.base.models.res_users': 9}
 {'coupling': 0.177583, 'cohesion': 0.469743, 'value': 51.5435, 'clusterSize': 173}
 {'coupling': 0.17644, 'cohesion': 0.467068, 'value': 51.5693, 'clusterSize': 174}
 {'coupling': 0.174457, 'cohesion': 0.462889, 'value': 51.764, 'clusterSize': 176}


 {'coupling': 0.344848, 'cohesion': 0.359244, 'value': 4.0233, 'clusterSize': 210}
{'coupling': 0.351066, 'cohesion': 0.355901, 'value': 2.0202, 'clusterSize': 211}
{'coupling': 0.358837, 'cohesion': 0.230358, 'value': -26.2377, 'clusterSize': 212}

{'coupling': 0.478486, 'cohesion': 0.113674, 'value': -93.4861, 'clusterSize': 259}
{'coupling': 0.466474, 'cohesion': 0.115346, 'value': -89.591, 'clusterSize': 258}
{'coupling': 0.477448, 'cohesion': 0.113507, 'value': -93.6245, 'clusterSize': 260


// tidak ada panggilan diri sendiri ???
selfCG['addons.web_tour']