package main

const blockMapping = `
{
  "settings": {
    "number_of_shards": 1,
    "number_of_replicas": 0
  },
  "mappings": {
    "block": {
      "properties": {
        "hash": {
          "type": "keyword"
        },
        "strippedsize": {
          "type": "integer"
        },
        "size": {
          "type": "integer"
        },
        "weight": {
          "type": "integer"
        },
        "height": {
          "type": "integer"
        },
        "versionHex": {
          "type": "text"
        },
        "merkleroot": {
          "type": "text"
        },
        "tx": {
          "properties": {
            "hex": {
              "type": "text"
            },
            "txid": {
              "type": "text"
            },
            "hash": {
              "type": "text"
            },
            "version": {
              "type": "short"
            },
            "size": {
              "type": "integer"
            },
            "vsize": {
              "type": "integer"
            },
            "locktime": {
              "type": "long"
            },
            "vin": {
              "properties": {
                "txid": {
                  "type": "text"
                },
                "vout": {
                  "type": "short"
                },
                "scriptSig": {
                  "properties": {
                    "asm": {
                      "type": "text"
                    },
                    "hex": {
                      "type": "text"
                    }
                  }
                },
                "sequence": {
                  "type": "long"
                },
                "txinwitness": {
                  "type":"keyword"
                }
              }
            },
            "vout": {
              "properties": {
                "value": {
                  "type": "double"
                },
                "n": {
                  "type": "short"
                },
                "scriptPubKey": {
                  "properties": {
                    "asm": {
                      "type": "text"
                    },
                    "hex": {
                      "type": "text"
                    },
                    "reqSigs": {
                      "type": "short"
                    },
                    "type": {
                      "type": "text"
                    },
                    "addresses": {
                      "type":"keyword"
                    }
                  }
                }
              }
            }
          }
        },
        "time": {
          "type": "long"
        },
        "mediantime": {
          "type": "long"
        },
        "nonce": {
          "type": "long"
        },
        "bits": {
          "type": "text"
        },
        "difficulty": {
          "type": "double"
        },
        "chainwork": {
          "type": "text"
        },
        "previoushash": {
          "type": "text"
        },
        "nexthash": {
          "type": "text"
        }
      }
    }
  }
}`

const txMapping = `
{
  "settings": {
    "number_of_shards": 1,
    "number_of_replicas": 0
  },
  "mappings": {
		"tx": {
      "properties": {
        "txid": {
          "type": "text"
        },
        "fee": {
          "type": "double"
        },
				"blockhash": {
					"type": "text"
				},
        "vins": {
          "type": "nested",
          "properties": {
            "address": {
              "type": "text"
            },
            "value": {
              "type": "double"
            }
          }
        },
        "vouts": {
          "type": "nested",
          "properties": {
            "address": {
              "type": "text"
            },
            "value": {
              "type": "double"
            }
          }
        },
        "time": {
          "type": "long"
        }
      }
    }
  }
}`

const voutMapping = `
{
  "settings": {
    "number_of_shards": 1,
    "number_of_replicas": 0
  },
  "mappings": {
		"vout": {
      "properties": {
        "txidbelongto": {
          "type": "text"
        },
        "value": {
          "type": "double"
        },
        "voutindex": {
          "type": "short"
        },
        "coinbase": {
          "type": "boolean"
        },
        "addresses": {
          "type":"keyword"
        },
				"time": {
					"type": "long"
				},
        "used": {
          "properties": {
            "txid": {
              "type": "text"
            },
            "vinindex": {
              "type": "short"
            }
          }
        }
      }
    }
  }
}`

const balanceMapping = `
{
  "settings": {
    "number_of_shards": 1,
    "number_of_replicas": 0
  },
  "mappings": {
		"balance": {
			"properties": {
				"address": {
					"type":"keyword"
				},
				"amount": {
					"type": "double"
				}
			}
		}
  }
}`