# SDOF - motion of Simgle Degree Of Freedom system.

1質点系の振動に関するパッケージです。

## Installation

```shell
> go install github.com/takatoh/SODF
```

## １質点系の自由振動

自由振動の基本的な関数は `freevibration` サブパッケージにまとめられています。

### 固有円振動数
```go
omega := freevibration.NaturalCircularFrequency(m, k)
```

### 固有振動数
```go
f := freevibration.NaturalFrequency(m, k)
```

### 固有周期
```go
t := freevibration.NaturalPeriod(m, k)
```

### 減衰固有円振動数
```go
omega_d := freevibration.dampedNaturalCircularFrequency(omega, h)
```

### 減衰固有周期
```go
t_d := freevibration.DampedNaturalPeriod(omega, h)
```

## １質点系の地震応答

直接積分法により応答時刻歴（絶対加速度、相対速度、相対変位）を求める関数を `directintegration` サブパッケージにまとめています。

- 平均加速度法
- 線形加速度法
- ニガムの方法
- ウィルソンのΘ法
- ルンゲ-クッタ法（RK4）

いずれの関数も次の5つの引数をとります。

- 減衰定数 `h`
- 固有振動数 `w`
- 時間間隔 `dt`
- データ数 `nn`
- 地震の地動加速度時刻歴 `ddy`

時間間隔は十分に小さいことを前提として、そのまま積分間隔としています。
一般に、地震の加速度時刻歴の時間間隔はそのまま積分するには大きいので、適切な方法で時間間隔を分割し、データを補間してください。

いずれの関数も次の3つの時刻歴を返します。

- 絶対加速度
- 相対速度
- 相対変位

### 平均加速度法
```go
acc, vel, dis := directintegration.AverageAcc(h, w, dt, nn, ddy)
```

### 線形加速度法
```go
acc, vel, dis := directintegration.LinearAcc(h, w, dt, nn, ddy)
```

### ニガムの方法
```go
acc, vel, dis := directintegration.NigamAcc(h, w, dt, nn, ddy)
```

### ウィルソンのΘ法
```go
acc, vel, dis := directintegration.WilsonTheta(h, w, dt, nn, ddy)
```

### （4次の）ルンゲ-クッタ法（RK4）
```go
acc, vel, dis := directintegration.RK4(h, w, dt, nn, ddy)
```
