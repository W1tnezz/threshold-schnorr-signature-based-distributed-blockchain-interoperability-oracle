
import faulthandler;
faulthandler.enable()

import matplotlib

matplotlib.use('Agg')
import matplotlib.pyplot as plt
from matplotlib import font_manager
import numpy as np
import pandas as pd


fig, ax = plt.subplots()

sig_nums = []

for i in range(100):
    if(i % 2 == 1):
        sig_nums.append(i)
    

ecdsa_verify_costs = []

for i in sig_nums:
    ecdsa_verify_costs.append(i * 4359)
bls_verify_costs = [113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969]

schnorr_verify_costs = [57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241, 57241]


plt.plot(sig_nums[0:25], ecdsa_verify_costs[0:25],  color='deeppink', label="ECDSA", marker = "*")
plt.plot(sig_nums[0:25], bls_verify_costs[0:25],  color='royalblue', label="Weighted BLS-multi", marker = "^")
plt.plot(sig_nums[0:25], schnorr_verify_costs[0:25],  color='limegreen', label="Weighted Schnorr-multi", marker = ".")


plt.gcf().subplots_adjust(left=0.15,top=0.9,bottom=0.1)
plt.xlabel("Total weight-value")  # 横坐标名字
plt.ylabel("Gas consumption")  # 纵坐标名字
plt.legend()
my_x_ticks = np.arange(0, 51, 5)
my_y_ticks = np.arange(0, 220001, 20000)
plt.xticks(my_x_ticks)
plt.yticks(my_y_ticks)
fig.savefig('签名验证gas消耗对比.svg', dpi=3200, format='svg')


#字典中的key值即为csv中列名
dataframe = pd.DataFrame({'signatures_number':sig_nums[0:25],'ecdsa':ecdsa_verify_costs[0:25],'BLS-multi':bls_verify_costs[0:25], 'Schnorr-multi':schnorr_verify_costs[0:25]})

#将DataFrame存储为csv,index表示是否显示行名，default=True
dataframe.to_csv("./csv/test.csv",sep=',')