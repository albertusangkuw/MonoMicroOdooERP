from odoo import api, fields, models, _

# Dynamic Call
class A(models.Model):
    _name = "aA"
    b = fields.Many2one('B')
    
class B(models.Model):
    _name = "aB"


class C(models.Model):
    _name = "aC"
    _inherit = ['A']

class D(models.Model):
    _inherit = 'aB'

# Static Call
class E(models.Model):
    _name = "aE"
    def defE():
        print("E")

class F(models.Model):
    _name = "aF"
    def caller():
        G.defG()
        E.defE()
        
class G(models.Model):
    _name = "aG"
    def defG():
        print("G")

