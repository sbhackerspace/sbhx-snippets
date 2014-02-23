#!/usr/bin/python -tt
#
# Copyright 2014 Garrett D Holmstrom
#
# This program is free software; you can redistribute it and/or modify
# it under the terms of the GNU General Public License as published by
# the Free Software Foundation; either version 3 of the License or (at
# your option) any later version accepted by the membership of the Santa
# Barbara Hackerspace (or its successor approved by the membership of
# the Santa Barbara Hackerspace), which shall act as a proxy as defined
# in Section 14 of version 3 of the license.
#
# This program is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU
# General Public License for more details.
#
# You should have received a copy of the GNU General Public License
# along with this program. If not, see http://www.gnu.org/licenses/.

from flask import Flask, jsonify, request
import subprocess


RADIUS_SERVER = 'radius.example.com'
RADIUS_SECRET = 'testing123'


app = Flask(__name__)


@app.route('/api/1/auth/basic', methods=['GET'])
def auth_basic():
    auth = request.authorization
    if not auth:
        return (jsonify(status='unauthorized'), 401,
                {'WWW-Authenticate': 'Basic realm="Login Required"'})
    else:
        username = auth.username.split('\n', 1)[0]
        password = auth.password.split('\n', 1)[0]
        code = check_radius_auth(username, password)
        if code == 200:
            return jsonify(status='ok'), 200
        elif code == 403:
            return jsonify(status='forbidden'), 403
        else:
            return jsonify(status='internal-server-error'), 500


def check_radius_auth(username, password):
    radclient = subprocess.Popen(
        ('radclient', RADIUS_SERVER, 'auth', RADIUS_SECRET),
        stdin=subprocess.PIPE, stdout=subprocess.PIPE)
    avps = {'User-Name': username,
            'User-Password': password}
    avps_str = '\n'.join('='.join((attr, val))
                         for attr, val in sorted(avps.iteritems()))
    output = radclient.communicate(avps_str)[0]
    # Received response ID 185, code 2, length = 20
    code = output.split('\n', 1)[0].partition(', code ')[2].split(',', 1)[0]
    if code == '2':
        # Access-Accept
        return 200
    elif code == '3':
        # Access-Reject
        return 403
    else:
        return 500


if __name__ == '__main__':
    # Don't forget to use HTTPS if using this in production.
    app.run(host='0.0.0.0')
