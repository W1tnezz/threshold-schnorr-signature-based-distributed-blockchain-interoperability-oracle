
import faulthandler;
faulthandler.enable()

import matplotlib

matplotlib.use('Agg')
import matplotlib.pyplot as plt
from matplotlib import font_manager
import numpy as np
import pandas as pd

def appendArr(arr, num, a, base):
    # cost = ax + b
    for i in range(num):
        if(i % 2 == 1):
            element = i * a + base
            arr.append(element)

sig_nums = []
ecdsa_verify_costs = []
bls_verify_costs = []
schnorr_verify_costs = []
bls_real_cost = []



ecdsa_base = 0
ecdsa_a = 4359

bls_base = 145192 + 15000
bls_a = 16666

schnorr_base = 57241
schnorr_a = 16666


appendArr(sig_nums, 100, 1, 0)

print(len(sig_nums))

appendArr(ecdsa_verify_costs, 100, ecdsa_a, ecdsa_base)

appendArr(bls_verify_costs, 100, bls_a, bls_base)

appendArr(schnorr_verify_costs, 100, schnorr_a, schnorr_base)

print(len(sig_nums))
print(ecdsa_verify_costs)
print(schnorr_verify_costs)


bls_threshold_verify_costs = [113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969]

schnorr_threshold_verify_costs = [57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241]




fig, ax = plt.subplots()
# font = font_manager.FontProperties(fname="/usr/share/fonts/truetype/ubuntu/UbuntuMono-B.ttf")

# plt.plot(sig_nums[2:41], ecdsa_verify_costs[2:41],  color='deepskyblue', label="ECDSA")
plt.plot(sig_nums[0:15], bls_threshold_verify_costs[0:15],  color='red', label="BLS-threshold", marker = "*")
plt.plot(sig_nums[0:15], schnorr_threshold_verify_costs[0:15],  color='green', label="Our scheme", marker = "^")
plt.plot(sig_nums[0:15], ecdsa_verify_costs[0:15],  color='blueviolet', label="ECDSA", marker = "*")
plt.plot(sig_nums[0:15], bls_verify_costs[0:15],  color='royalblue', label="BLS-multi", marker = "^")
plt.plot(sig_nums[0:15], schnorr_verify_costs[0:15],  color='orange', label="Schnorr-multi", marker = ".")


plt.gcf().subplots_adjust(left=0.15,top=0.9,bottom=0.1)
plt.xlabel("Number of signatures")  # 横坐标名字
plt.ylabel("Gas consumption")  # 纵坐标名字
plt.legend()
my_x_ticks = np.arange(0, 31, 5)
my_y_ticks = np.arange(0, 700001, 50000)
plt.xticks(my_x_ticks)
plt.yticks(my_y_ticks)
plt.grid()
fig.savefig('./figures/链上聚合公钥签名验证gas消耗对比.svg', dpi=3200, format='svg')

#字典中的key值即为csv中列名
dataframe = pd.DataFrame({'signatures_number':sig_nums[0:15],'BLS-multi':bls_verify_costs[0:15], 'Schnorr-multi':schnorr_verify_costs[0:15], 'BLS-Threshold': bls_threshold_verify_costs[0:15],'Schnorr-Threshold': schnorr_threshold_verify_costs[0:15]})

#将DataFrame存储为csv,index表示是否显示行名，default=True
dataframe.to_csv("./csv/combine+verify.csv",sep=',')



