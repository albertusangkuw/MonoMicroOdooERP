from odoo.http import request, SESSION_LIFETIME
import jwt
import datetime


class JWTHandler:
    _algorithm = "HS256"
    _secret = "GyIXT3pB0pVGK68kQPFG0QCrpOq6033A"

    def _get_meta_user_by_uid(self, uid):
        user = request.env['res.users'].sudo().search([('id', '=', uid)], limit=1)
        if user:
            current_time = datetime.datetime.utcnow()
            expiration_time = current_time + datetime.timedelta(seconds=SESSION_LIFETIME)
            exp = int(expiration_time.timestamp())
            return {
                'uid': user.id,
                'user_context': request.env.context,
                'db': request.env.cr.dbname,
                'company_id': user.company_id.id,
                'partner_id': user.partner_id.id,
                'group_id': user.groups_id.ids,
                'exp' :exp,
            }
        else:
            return {
                'uid': user.id,
                'user_context': {},
                'db': '',
                'company_id': '',
                'partner_id': '',
                'group_id' : 0,
                'exp': 0,
            }

    def encode(self, uid):
        return jwt.encode(self._get_meta_user_by_uid(uid), self._secret, algorithm=self._algorithm)
