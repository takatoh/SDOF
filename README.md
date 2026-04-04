# SDOF - motion of Simgle Degree Of Freedom system.

1質点系の振動に関するパッケージです。

## Installation

```
> go install github.com/takatoh/SODF
```

## １質点系の自由振動

自由振動の基本的な関数は `freevibration` サブパッケージにまとめられています。

### 固有円振動数
```
omega := freevibration.NaturalCircularFrequency(m, k)
```

### 固有振動数
```
f := freevibration.NaturalFrequency(m, k)
```

### 固有周期
```
t := freevibration.NaturalPeriod(m, k)
```

### 減衰固有円振動数
```
omega_d := freevibration.dampedNaturalCircularFrequency(omega, h)
```

### 減衰固有周期
```
t_d := freevibration.DampedNaturalPeriod(omega, h)
```
