"""
    Ais-Stream WebsocketObjects

    A sample API to illustrate OpenAPI concepts  # noqa: E501

    The version of the OpenAPI document: 1.0.0
    Generated by: https://openapi-generator.tech
"""


import sys
import unittest

import aisStream
from aisStream.model.channel_management_area import ChannelManagementArea
from aisStream.model.channel_management_unicast import ChannelManagementUnicast
globals()['ChannelManagementArea'] = ChannelManagementArea
globals()['ChannelManagementUnicast'] = ChannelManagementUnicast
from aisStream.model.channel_management import ChannelManagement


class TestChannelManagement(unittest.TestCase):
    """ChannelManagement unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def testChannelManagement(self):
        """Test ChannelManagement"""
        # FIXME: construct object with mandatory attributes with example values
        # model = ChannelManagement()  # noqa: E501
        pass


if __name__ == '__main__':
    unittest.main()
