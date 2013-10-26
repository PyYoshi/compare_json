#!/usr/bin/env python
# coding: utf-8

import json
import simplejson
import ujson
import time

def load_json():
	return open('./public_timeline.json').read()

def decode(m, json, loop):
	print('=== %s decode ===' % m.__name__)
	total_time = 0
	for i in range(loop):
		start = time.time()
		m.loads(json)
		c = time.time() - start
		total_time += c
	print('%s call(s)/s' % (1/(total_time/loop)))

def encode(m, obj, loop):
	print('=== %s encode ===' % m.__name__)
	total_time = 0
	for i in range(loop):
		start = time.time()
		m.dumps(obj)
		c = time.time() - start
		total_time += c
	print('%s call(s)/s' % (1/(total_time/loop)))

def main():
	loop = 10000
	json_string = load_json()
	json_obj = json.loads(json_string)
	modules = [json, simplejson, ujson]
	for module in modules:
		decode(module, json_string, loop)
		encode(module, json_obj, loop)

if '__main__' in __name__:
	main()